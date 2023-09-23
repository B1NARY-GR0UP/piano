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
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sync"

	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

type Handler interface {
	ServeHTTP(ctx context.Context, pk *PianoKey)
}

// HandlerFunc is the core type of PIANO
type HandlerFunc func(ctx context.Context, pk *PianoKey)

// HandlersChain is the slice of HandlerFunc
type HandlersChain []HandlerFunc

// PianoKey play the piano with PianoKeys
type PianoKey struct {
	Request *http.Request
	Writer  http.ResponseWriter

	index    int // initialize with -1
	Params   Params
	handlers HandlersChain
	rwMutex  sync.RWMutex
	KVs      M
}

const (
	breakIndex = math.MaxInt8 / 2
)

// NewContext will return a new context object which is piano key
func NewContext(maxParams uint16) *PianoKey {
	ps := make(Params, 0, maxParams)
	return &PianoKey{
		Params: ps,
		index:  -1,
	}
}

// SetHandlers will set handlers field for PianoKey context
func (pk *PianoKey) SetHandlers(handlers HandlersChain) {
	pk.handlers = handlers
}

// Next executes the handlers on the chain
func (pk *PianoKey) Next(ctx context.Context) {
	pk.index++
	for pk.index < len(pk.handlers) {
		pk.handlers[pk.index](ctx, pk)
		pk.index++
	}
}

// Set will store the key and value into this PianoKey
func (pk *PianoKey) Set(key string, value any) {
	pk.rwMutex.Lock()
	// lazy initializes
	if pk.KVs == nil {
		pk.KVs = make(M)
	}
	pk.KVs[key] = value
	pk.rwMutex.Unlock()
}

// Get will return the value corresponding to the given key, it will return (nil, false) if key does not exist
func (pk *PianoKey) Get(key string) (value any, ok bool) {
	pk.rwMutex.RLock()
	value, ok = pk.KVs[key]
	pk.rwMutex.RUnlock()
	return
}

// MustGet will return the value corresponding to the given key, it will panic if key does not exist
func (pk *PianoKey) MustGet(key string) any {
	if value, ok := pk.Get(key); ok {
		return value
	}
	panic("key \"" + key + "\" does not exist")
}

// Param return the corresponding param in request URL
func (pk *PianoKey) Param(key string) string {
	return pk.Params.Get(key)
}

// Break current handler
func (pk *PianoKey) Break() {
	pk.index = breakIndex
}

func (pk *PianoKey) BreakWithStatus(code int) {
	pk.SetStatusCode(code)
	pk.Break()
}

func (pk *PianoKey) BreakWithMessage(code int, msg string) {
	pk.SetStatusCode(code)
	pk.SetHeader(consts.HeaderContentType, "text/plain; charset=utf-8")
	_, _ = fmt.Fprintf(pk.Writer, msg)
	pk.Break()
}

// Query is used to match HTTP GET query params
func (pk *PianoKey) Query(key string) string {
	return pk.Request.URL.Query().Get(key)
}

// DefaultQuery is Query with default value when no match
func (pk *PianoKey) DefaultQuery(key, defaultValue string) string {
	value := pk.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// PostForm is used to get HTTP POST form data
func (pk *PianoKey) PostForm(key string) string {
	return pk.Request.PostFormValue(key)
}

// DefaultPostForm is PostForm with default value when no match
func (pk *PianoKey) DefaultPostForm(key, defaultValue string) string {
	value := pk.PostForm(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (pk *PianoKey) FormValue(key string) string {
	return pk.Request.FormValue(key)
}

func (pk *PianoKey) DefaultFormValue(key, defaultValue string) string {
	value := pk.FormValue(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// SetStatusCode is used to set HTTP response code
func (pk *PianoKey) SetStatusCode(code int) {
	pk.Writer.WriteHeader(code)
}

// SetHeader is used to set HTTP response header
func (pk *PianoKey) SetHeader(key, value string) {
	if value == "" {
		pk.Writer.Header().Del(key)
		return
	}
	pk.Writer.Header().Set(key, value)
}

func (pk *PianoKey) writeJSON(data any) error {
	pk.SetHeader(consts.HeaderContentType, "application/json; charset=utf-8")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = pk.Writer.Write(jsonBytes)
	if err != nil {
		return err
	}
	return nil
}

// JSON is used to response data in JSON form
func (pk *PianoKey) JSON(code int, data any) {
	pk.SetStatusCode(code)
	err := pk.writeJSON(data)
	if err != nil {
		panic(err)
	}
}

func (pk *PianoKey) writeString(format string, data ...any) error {
	pk.SetHeader(consts.HeaderContentType, "text/plain; charset=utf-8")
	// Fprintf will pass the data to the writer
	_, err := fmt.Fprintf(pk.Writer, format, data...)
	return err
}

// String is used to response data in string form
func (pk *PianoKey) String(code int, format string, data ...any) {
	pk.SetStatusCode(code)
	err := pk.writeString(format, data...)
	if err != nil {
		panic(err)
	}
}

// refresh will reset the PianoKey except the request, writer and params field
func (pk *PianoKey) refresh() {
	pk.index = -1
	pk.handlers = nil
	pk.KVs = nil
}
