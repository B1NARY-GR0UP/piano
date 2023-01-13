package core

import (
	"net/http"
	"strings"
)

// RouterGroup must implement IRouter
var _ IRouter = (*RouterGroup)(nil)

type IRouter interface {
	IRoute
	GROUP(string, ...HandlerFunc) *RouterGroup
}

type IRoute interface {
	USE(...HandlerFunc)

	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
}

type RouterGroup struct {
	engine   *Engine
	Handlers HandlersChain
	basePath string
	isRoot   bool
}

// GROUP will new a route group
func (rg *RouterGroup) GROUP(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: rg.combineHandlers(handlers),
		basePath: rg.calculateAbsolutePath(relativePath),
		engine:   rg.engine,
	}
}

// USE middleware or other custom handlers
func (rg *RouterGroup) USE(middleware ...HandlerFunc) {
	rg.Handlers = append(rg.Handlers, middleware...)
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
