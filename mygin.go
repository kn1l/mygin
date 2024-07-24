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
	// handlers
	for context.index < len(context.handlers) {
		context.Next()
	}
	// fmt.Fprintf(w, "\ndebug: %s\n", context.Path)
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
	node.setHandlers(handlers)
}

func (engine *Engine) Group(absolutePath string, handlers ...HandlerFunc) *RouterGroup {
	return nil
}

func (engine *Engine) Handle(method, path string, handlers ...HandlerFunc) {
	engine.addRoute(method, path, handlers)
}

func (engine *Engine) GET(path string, handlers ...HandlerFunc) {
	engine.Handle(http.MethodGet, path, handlers...)
}

func (engine *Engine) Run(addr ...string) error {
	address := resolveAddress(addr...)
	fmt.Printf("address: %s\n", address)
	return http.ListenAndServe(address, engine)
}
