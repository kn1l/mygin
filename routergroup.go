package mygin

import "net/http"

type RouterGroup struct {
	Handlers HandlerFuncChain
	fullPath string
	engine   *Engine
}

func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	absolutePath := group.calcAbsolutePath(relativePath)
	mergedHandlers := append(group.Handlers, handlers...)
	return group.engine.Group(absolutePath, mergedHandlers...)
}

// func (group *RouterGroup) addRoute(method, path string, handlers HandlerFuncChain) {
// 	root := group.engine.trees.getMethodTree(method)
// 	node := root.insert(path)
// 	node.handlers = handlers
// }

func (group *RouterGroup) Handle(method, relativePath string, handlers ...HandlerFunc) {
	absolutePath := group.calcAbsolutePath(relativePath)
	println(absolutePath)
	group.engine.addRoute(method, absolutePath, handlers)
}

func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	group.Handle(http.MethodGet, relativePath, handlers...)
}

func (group *RouterGroup) calcAbsolutePath(relativePath string) string {
	var absolutePath string
	if relativePath == "/" {
		absolutePath = group.fullPath
	} else if group.fullPath == "/" {
		absolutePath = relativePath
	} else {
		absolutePath = group.fullPath + relativePath
	}
	return absolutePath
}
