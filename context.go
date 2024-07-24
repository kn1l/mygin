package mygin

import (
	"fmt"
	"net/http"
)

type Param struct {
	Key   string
	Value string
}

type Params []Param

type Context struct {
	// origin objects
	Writer  http.ResponseWriter
	Request *http.Request

	handlers HandlerFuncChain
	index    int
	Params   Params

	Method   string
	Pathlist []string
	engine   *Engine
}

func (engine *Engine) newContext(w http.ResponseWriter, req *http.Request) *Context {
	c := &Context{
		Writer:  w,
		Request: req,
		Method:  req.Method,
		engine:  engine,
	}
	c.Pathlist = splitPath(req.URL.Path)

	c.setHandlers()
	return c
}

func (c *Context) setHandlers() {

	root := c.engine.trees.getMethodTree(c.Method)

	if node := root.search(c, 0); node != nil && node.handlers != nil {
		c.handlers = node.handlers
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 not Found !!!")
		})
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

func (c *Context) String(statuscode int, format string, values ...any) {
	c.Writer.WriteHeader(statuscode)
	fmt.Fprintf(c.Writer, format, values...)
}

func (c *Context) Html() {

}
