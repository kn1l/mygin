package mygin

type RouterGroup struct {
	Handlers HandlerFuncChain
	Path     string
	engine   *Engine
}

func newRouterGroup(path string) *RouterGroup {
	routers := &RouterGroup{}
	return routers
}

func (group *RouterGroup) Group(path string) *RouterGroup {
	return group
}

func (routers *RouterGroup) Handle(method, relativePath string, handlers ...HandlerFunc) {
	root := routers.engine.trees.getMethodTree(method)
	node := root.insert(relativePath)
	node.setHandlers(handlers)
}

func (routers *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	routers.Handle("GET", relativePath, handlers...)
}
