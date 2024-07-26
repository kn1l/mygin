package mygin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// map[string]any
type H map[string]interface{}

type Param struct {
	Key   string
	Value string
}

type Params []Param

type Context struct {
	// origin objects
	Writer  ResponseWriter
	Request *http.Request

	handlers HandlerFuncChain
	index    int
	Params   Params

	Method   string
	Path     string
	Pathlist []string
	engine   *Engine
}

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (engine *Engine) newContext(w http.ResponseWriter, req *http.Request) *Context {
	c := &Context{
		Writer:  ResponseWriter{ResponseWriter: w},
		Request: req,
		Method:  req.Method,
		Path:    req.URL.Path,
		engine:  engine,
	}
	c.Pathlist = splitPath(req.URL.Path)
	c.Path = joinPath(c.Pathlist)
	c.setHandlers()
	return c
}

func (c *Context) setHandlers() {
	root := c.engine.trees.getMethodTree(c.Method)
	if node := root.search(c, 0); node != nil && node.handlers != nil {
		c.handlers = node.handlers
	} else {
		mergedHandlers := append(c.engine.Handlers, func(c *Context) {
			notFoundHandler(c)
		})
		c.handlers = c.engine.addRoute(c.Method, c.Path, mergedHandlers).handlers
	}
}

func (c *Context) Next() {
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	} else {
		errorPrint("invalid handlers index!")
	}
}

func (c *Context) Param(key string) string {
	for _, p := range c.Params {
		if p.Key == key {
			return p.Value
		}
	}
	return ""
}

func (c *Context) String(code int, format string, values ...any) {
	c.Writer.WriteHeader(code)
	c.Writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(c.Writer, format, values...)
}

func (c *Context) JSON(code int, obj any) {
	c.Writer.WriteHeader(code)
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(obj)
}

func (c *Context) HTML(code int, name string, obj any) {
	c.Writer.WriteHeader(code)
	c.Writer.Header().Set("Content-Type", "text/html")
	c.engine.htmlTemplate.Lookup(name).Execute(c.Writer, obj)
}

func notFoundHandler(c *Context) {
	c.Writer.WriteHeader(http.StatusNotFound)
	c.Writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(c.Writer, "404 not Found !!!")
}
