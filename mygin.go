package mygin

import (
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
}

// New returns an instance of *Engine
func New() *Engine {
	engine := &Engine{}
	return engine
}

func (engine *Engine) addRoute(method, path string, handler HandlerFuncChain) {

}

func (engine *Engine) Handle(method, relativePath string, handlers ...HandlerFunc) {
	root := engine.trees.getMethodTree(method)
	node := root.insert(relativePath)
	node.setHandlers(handlers)
}

func (engine *Engine) GET(relativePath string, handlers ...HandlerFunc) {
	engine.Handle("GET", relativePath, handlers...)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
