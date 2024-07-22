package mygin

import (
	"fmt"
	"net/http"
	"strings"
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

	Method string
	Path   string
	engine *Engine
}

func (engine *Engine) newContext(w http.ResponseWriter, req *http.Request) *Context {
	c := &Context{
		Writer:  w,
		Request: req,
		Method:  req.Method,
		Path:    req.URL.Path,
		engine:  engine,
	}
	c.setHandlers()
	return c
}

func (c *Context) setHandlers() {
	pathlist := make([]string, 0)
	for _, p := range strings.Split(c.Path, "/") {
		if p != "" {
			pathlist = append(pathlist, p)
		}
	}
	root := c.engine.trees.getMethodTree(c.Method)
	n := root
	for _, p := range pathlist {
		subpath := "/" + p
		isFound := false
		for _, child := range n.children {
			if child.path == subpath {
				isFound = true
				n = child
				break
			}
		}
		if !isFound {
			c.handlers = append(c.handlers, func(c *Context) {
				c.String(http.StatusNotFound, "404 not Found !!!")
			})
		}
	}

	// if node := root.search(c.Path); node != nil {
	// 	c.handlers = node.handlers
	// } else {
	// 	c.handlers = append(c.handlers, func(c *Context) {
	// 		c.String(http.StatusNotFound, "404 not Found !!!")
	// 	})
	// }
}

func (c *Context) setParams() {

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
	return ""
}

func (c *Context) String(statuscode int, format string, values ...any) {
	c.Writer.WriteHeader(statuscode)
	fmt.Fprintf(c.Writer, format, values...)
}

func (c *Context) Html() {

}
