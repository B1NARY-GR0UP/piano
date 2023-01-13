package core

import (
	"context"
	"net/http"
	"strings"
	"sync"

	"github.com/B1NARY-GR0UP/inquisitor/core"
)

type Engine struct {
	// RouterGroup is a composition of Engine so that Engine can use RouterGroup functions
	RouterGroup

	forest    MethodForest
	options   *Options
	ctxPool   sync.Pool
	maxParams uint16
}

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

// Play Start the Server
func (e *Engine) Play() {
	core.Infof("PIANO server is listening on address %v", e.options.Addr)
	err := http.ListenAndServe(e.options.Addr, e)
	if err != nil {
		panic("PIANO Server Start Failed")
	}
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

// handleRequest will handle HTTP request
func (e *Engine) handleRequest(ctx context.Context, pk *PianoKey) {
	matchedTree, ok := e.forest.get(pk.Request.Method)
	if !ok {
		e.serveError(ctx, pk)
	}
	path := pk.Request.URL.Path
	validatePath(path)
	// note: must use pointer
	params := &pk.Params
	matchedNode := matchedTree.search(path, params)
	if matchedNode == nil {
		e.serveError(ctx, pk)
	}
	pk.SetHandlers(matchedNode.handlers)
	pk.Next(ctx)
}

func (e *Engine) serveError(_ context.Context, pk *PianoKey) {
	pk.handlers = append(pk.handlers, func(ctx context.Context, pk *PianoKey) {
		pk.String(http.StatusNotFound, "404 Page Not Found")
	})
}

// addRoute will add path into trie tree
func (e *Engine) addRoute(method, path string, handlers HandlersChain) {
	isValid := validateRoute(method, path, handlers)
	if !isValid {
		panic("please check your route")
	}
	core.Infof("Register route: [%v] %v", strings.ToUpper(method), path)
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
