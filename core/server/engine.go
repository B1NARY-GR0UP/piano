// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package server

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/B1NARY-GR0UP/inquisitor/core"
	"github.com/B1NARY-GR0UP/piano/pkg/bytesconv"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
	"github.com/B1NARY-GR0UP/piano/pkg/errors"
	"github.com/B1NARY-GR0UP/piano/pkg/nocopy"
)

type Engine struct {
	_ nocopy.NoCopy

	RouterGroup

	// basic
	forest    MethodForest
	options   *Options
	ctxPool   sync.Pool
	maxParams uint16
	// initialized | running | shutdown | closed
	status uint32

	// hook
	OnRun      []HookFuncWithErr
	OnShutdown []HookFunc

	// TODO: support transport
}

type (
	HookFunc        func(ctx context.Context)
	HookFuncWithErr func(ctx context.Context) error
)

const (
	_ uint32 = iota
	statusInitialized
	statusRunning
	statusShutdown
	statusClosed
)

var (
	errAlreadyInit    = errors.NewPrivate("engine has been init already")
	errAlreadyRunning = errors.NewPrivate("engine is already running")
	errNotRunning     = errors.NewPrivate("engine is not running")
	errShutdown       = errors.NewPrivate("engine shutdown error")
)

// NewEngine for PIANO
func NewEngine(opts *Options) *Engine {
	e := &Engine{
		forest: make(MethodForest, 0, 5),
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			isRoot:   true,
		},
		options: opts,
	}
	e.RouterGroup.engine = e
	// TODO: fix maxParam (assigned?)
	e.ctxPool.New = func() any {
		return e.allocateContext(e.maxParams)
	}
	return e
}

// Init PIANO engine
func (e *Engine) Init() error {
	if !atomic.CompareAndSwapUint32(&e.status, 0, statusInitialized) {
		return errAlreadyInit
	}
	return nil
}

// Run Start the PIANO Engine
func (e *Engine) Run() error {
	if err := e.Init(); err != nil {
		return err
	}
	if !atomic.CompareAndSwapUint32(&e.status, statusInitialized, statusRunning) {
		return errAlreadyRunning
	}
	defer atomic.StoreUint32(&e.status, statusClosed)
	if err := e.executeOnRunHooks(context.Background()); err != nil {
		return err
	}
	core.Infof("---PIANO--- Server is listening on address %v", e.options.Addr)
	return http.ListenAndServe(e.options.Addr, e)
}

func (e *Engine) Shutdown(ctx context.Context) error {
	if atomic.LoadUint32(&e.status) != statusRunning {
		return errNotRunning
	}
	if !atomic.CompareAndSwapUint32(&e.status, statusRunning, statusShutdown) {
		return errShutdown
	}
	ch := make(chan struct{})
	go e.executeOnShutdownHooks(ctx, ch)
	defer func() {
		select {
		case <-ctx.Done():
			core.Errorf("---PIANO--- Execute shutdown hooks timeout: %v", ctx.Err())
			return
		case <-ch:
			core.Info("---PIANO--- Execute shutdown hooks done")
			return
		}
	}()
	return nil
}

// ServeHTTP core function, replace DefaultServeMux
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pk := e.ctxPool.Get().(*PianoKey)
	// inject content to *PianoKey
	pk.Writer = w
	pk.Request = req
	pk.refresh()
	e.handleRequest(context.Background(), pk)
	e.ctxPool.Put(pk)
}

func (e *Engine) IsRunning() bool {
	return atomic.LoadUint32(&e.status) == statusRunning
}

// Options return options field of current engine
func (e *Engine) Options() *Options {
	return e.options
}

// handleRequest will handle HTTP request
func (e *Engine) handleRequest(ctx context.Context, pk *PianoKey) {
	matchedTree, ok := e.forest.get(pk.Request.Method)
	if !ok {
		e.serveError(ctx, pk, consts.StatusMethodNotAllowed, bytesconv.S2B(consts.BodyMethodNotAllowed))
	}
	path := pk.Request.URL.Path
	validatePath(path)
	// note: must use pointer
	params := &pk.Params
	matchedNode := matchedTree.search(path, params)
	if matchedNode == nil {
		e.serveError(ctx, pk, consts.StatusNotFound, bytesconv.S2B(consts.BodyNotFound))
	}
	pk.SetHandlers(matchedNode.handlers)
	pk.Next(ctx)
}

func (e *Engine) serveError(ctx context.Context, pk *PianoKey, code int, message []byte) {
	pk.SetStatusCode(code)
	_, _ = pk.Writer.Write(message)
	pk.Next(ctx)
	// TODO: optimize after pack protocol
}

// addRoute will add path into trie tree
func (e *Engine) addRoute(method, path string, handlers HandlersChain) {
	isValid := validateRoute(method, path, handlers)
	if !isValid {
		core.Warnf("---PIANO--- Route %v is invalid", path)
	}
	core.Infof("---PIANO--- Register route: [%v] %v", strings.ToUpper(method), path)
	methodTree, ok := e.forest.get(method)
	// create a new method tree if no match in the forest
	if !ok {
		methodTree = &tree{
			method: method,
		}
		e.forest = append(e.forest, methodTree)
	}
	methodTree.addRoute(path, handlers)
	// update maxParams
	if paramCount := calculateParam(path); paramCount > e.maxParams {
		e.maxParams = paramCount
	}
}

func (e *Engine) allocateContext(maxParams uint16) *PianoKey {
	return NewContext(maxParams)
}

func (e *Engine) executeOnRunHooks(ctx context.Context) error {
	for _, h := range e.OnRun {
		if err := h(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) executeOnShutdownHooks(ctx context.Context, ch chan struct{}) {
	wg := sync.WaitGroup{}
	for _, h := range e.OnShutdown {
		wg.Add(1)
		go func(hook HookFunc) {
			defer wg.Done()
			hook(ctx)
		}(h)
	}
	wg.Wait()
	ch <- struct{}{}
}
