package mygin

import (
	"net/http"
	"text/template"
)

type HandlerFunc func(*Context)

type HandlerFuncChain []HandlerFunc

// Engine implements the interface Handler in net/http
type Engine struct {
	RouterGroup

	trees        methodTrees
	htmlTemplate *template.Template
}

// implements Handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := engine.newContext(w, req)
	// handlers
	for context.index < len(context.handlers) {
		context.Next()
	}
	requestInfoPrint(context.Writer.statusCode, req.Method, req.URL.Path)
}

func (engine *Engine) Run(addr ...string) error {
	address := resolveAddress(addr...)
	debugPrint("Listening and serving HTTP on %s", address)
	return http.ListenAndServe(address, engine)
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

// addRoute registers a route with its method,path and handlers
func (engine *Engine) addRoute(method, path string, handlers HandlerFuncChain) *node {
	root := engine.trees.getMethodTree(method)
	node := root.insert(path)
	node.handlers = handlers
	return node
}

func (engine *Engine) LoadHTMLFiles(filenames ...string) {
	engine.htmlTemplate = template.Must(template.ParseFiles(filenames...))
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplate = template.Must(template.ParseGlob(pattern))
}
