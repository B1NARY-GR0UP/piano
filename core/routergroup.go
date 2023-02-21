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

package core

import (
	"net/http"
	"regexp"
	"strings"
)

// RouterGroup must implement IRouter
var _ IRouter = (*RouterGroup)(nil)

type IRouter interface {
	IRoute
	Group(string, ...HandlerFunc) *RouterGroup
	Use(...HandlerFunc)
}

type IRoute interface {
	Handle(string, string, ...HandlerFunc)

	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)
}

type RouterGroup struct {
	engine   *Engine
	Handlers HandlersChain
	basePath string
	isRoot   bool
}

// Group will new a route group
func (rg *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: rg.combineHandlers(handlers),
		basePath: rg.calculateAbsolutePath(relativePath),
		engine:   rg.engine,
	}
}

// Use middlewares or other custom handlers
func (rg *RouterGroup) Use(middleware ...HandlerFunc) {
	rg.Handlers = append(rg.Handlers, middleware...)
}

// Handle is suggested to use for custom methods
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	if isMatch := regexp.MustCompile("^[A-Z]+$").MatchString(httpMethod); !isMatch {
		panic("http method " + httpMethod + " is not valid")
	}
	rg.handle(httpMethod, relativePath, handlers)
}

// GET will handler HTTP GET request
func (rg *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodGet, relativePath, handlers)
}

// POST will handler HTTP POST request
func (rg *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodPost, relativePath, handlers)
}

// PUT will handler HTTP PUT request
func (rg *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodPut, relativePath, handlers)
}

// DELETE will handler HTTP DELETE request
func (rg *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodDelete, relativePath, handlers)
}

// PATCH will handler HTTP PATCH request
func (rg *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodPatch, relativePath, handlers)
}

// OPTIONS will handler HTTP OPTIONS request
func (rg *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodOptions, relativePath, handlers)
}

// HEAD will handler HTTP HEAD request
func (rg *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) {
	rg.handle(http.MethodHead, relativePath, handlers)
}

// handle will handle HTTP request
func (rg *RouterGroup) handle(method, relativePath string, handlers HandlersChain) {
	absolutePath := rg.calculateAbsolutePath(relativePath)
	mergedHandlers := rg.combineHandlers(handlers)
	// note that the relative path is changed to absolutePath and handlers are changed to mergedHandlers
	rg.engine.addRoute(method, absolutePath, mergedHandlers)
}

// calculateAbsolutePath e.g. /binary + /lorain = /binary/lorain
func (rg *RouterGroup) calculateAbsolutePath(relativePath string) string {
	if relativePath == "" {
		return rg.basePath
	}
	sb := &strings.Builder{}
	if rg.basePath != "/" {
		sb.WriteString(rg.basePath)
	}
	sb.WriteString(relativePath)
	return sb.String()
}

// combineHandlers in different place
func (rg *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	totalSize := len(rg.Handlers) + len(handlers)
	mergedHandlers := make(HandlersChain, totalSize)
	copy(mergedHandlers, rg.Handlers)
	copy(mergedHandlers[len(rg.Handlers):], handlers)
	return mergedHandlers
}
