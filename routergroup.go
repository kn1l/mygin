package mygin

import "net/http"

type RouterGroup struct {
	Handlers HandlerFuncChain
	fullPath string
	engine   *Engine
}

func (engine *Engine) newRouterGroup(path string) *RouterGroup {
	group := &RouterGroup{
		fullPath: path,
	}
	return group
}

func (group *RouterGroup) Group(relativePath string) *RouterGroup {
	return group
}

func (group *RouterGroup) Handle(method, relativePath string, handlers ...HandlerFunc) {

}

func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	group.Handle(http.MethodGet, relativePath, handlers...)
}
