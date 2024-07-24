package mygin

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(*Context)

type HandlerFuncChain []HandlerFunc

// Engine implements the interface Handler in net/http
type Engine struct {
	RouterGroup

	trees methodTrees
}

// implements Handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := engine.newContext(w, req)
	println(len(context.handlers))
	// handlers
	for context.index < len(context.handlers) {
		context.Next()
	}
}

// New returns an instance of *Engine
func New() *Engine {
	engine := &Engine{}
	group := RouterGroup{
		fullPath: "/",
		engine:   engine,
	}
	engine.RouterGroup = group
	return engine
}

func Default() *Engine {
	engine := New()
	return engine
}

func (engine *Engine) addRoute(method, path string, handlers HandlerFuncChain) {
	root := engine.trees.getMethodTree(method)
	node := root.insert(path)
	node.handlers = handlers
}

func (engine *Engine) Group(absolutePath string, handlers ...HandlerFunc) *RouterGroup {
	if absolutePath == "/" {
		return &engine.RouterGroup
	}
	group := &RouterGroup{
		fullPath: absolutePath,
		engine:   engine,
	}
	group.Handlers = append(engine.RouterGroup.Handlers, handlers...)
	return group
}

func (engine *Engine) Use(handlers ...HandlerFunc) {
	engine.RouterGroup.Handlers = append(engine.RouterGroup.Handlers, handlers...)
}

func (engine *Engine) Handle(method, path string, handlers ...HandlerFunc) {
	mergedHandlers := append(engine.RouterGroup.Handlers, handlers...)
	engine.addRoute(method, path, mergedHandlers)
}

func (engine *Engine) GET(path string, handlers ...HandlerFunc) {
	engine.Handle(http.MethodGet, path, handlers...)
}

func (engine *Engine) POST(path string, handlers ...HandlerFunc) {
	engine.Handle(http.MethodPost, path, handlers...)
}

func (engine *Engine) PUT(path string, handlers ...HandlerFunc) {
	engine.Handle(http.MethodPut, path, handlers...)
}

func (engine *Engine) DELETE(path string, handlers ...HandlerFunc) {
	engine.Handle(http.MethodDelete, path, handlers...)
}

func (engine *Engine) HEAD(path string, handlers ...HandlerFunc) {
	engine.Handle(http.MethodHead, path, handlers...)
}

func (engine *Engine) OPTIONS(path string, handlers ...HandlerFunc) {
	engine.Handle(http.MethodOptions, path, handlers...)
}

func (engine *Engine) PATCH(path string, handlers ...HandlerFunc) {
	engine.Handle(http.MethodPatch, path, handlers...)
}

func (engine *Engine) Run(addr ...string) error {
	address := resolveAddress(addr...)
	fmt.Printf("address: %s\n", address)
	return http.ListenAndServe(address, engine)
}
