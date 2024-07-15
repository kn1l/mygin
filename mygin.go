package mygin

import (
	"fmt"
	"net/http"
)

// The HandlerFunc type is
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implements the interface Handler in net/http
type Engine struct {
	Router map[string]HandlerFunc
}

// implements Handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "404 Not Found")
}

// New returns an instance of *Engine
func New() *Engine {
	engine := &Engine{
		Router: make(map[string]HandlerFunc),
	}
	return engine
}

// GET
func GET(route string, f HandlerFunc) {

}

// TODO
func Handle() {

}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
