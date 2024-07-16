package mygin

type RouterGroup struct {
	Handlers HandlerFuncChain
	Path     string
	engine   *Engine
}
