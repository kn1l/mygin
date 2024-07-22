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
	return engine
}

func Default() *Engine {
	return New()
}

func (engine *Engine) addRoute(method, path string, handler HandlerFuncChain) {

}

func (engine *Engine) Group(path string, handlers ...HandlerFunc) *RouterGroup {
	// routers := engine
	return nil
}

func (engine *Engine) Handle(method, relativePath string, handlers ...HandlerFunc) {
	root := engine.trees.getMethodTree(method)
	node := root.insert(relativePath)
	node.setHandlers(handlers)
}

func (engine *Engine) GET(relativePath string, handlers ...HandlerFunc) {
	engine.Handle("GET", relativePath, handlers...)
}

func (engine *Engine) Run(addr ...string) error {
	address := resolveAddress(addr...)
	fmt.Printf("address: %s\n", address)
	return http.ListenAndServe(address, engine)
}
