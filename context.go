package mygin

import (
	"fmt"
	"net/http"
)

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path string

	// response info
	StatusCode int
}

func newContext() *Context {
	c := &Context{}
	return c
}

func (c *Context) String(statuscode int, format string, values ...any) {
	fmt.Fprintf(c.Writer, format, values...)
}

func (c *Context) Html() {

}
