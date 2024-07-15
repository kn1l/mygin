package mygin

import "net/http"

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path string

	// response info
	StatusCode int
}
