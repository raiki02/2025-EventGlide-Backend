package router

type PostRouterHdl interface {
	RegisterPostRouters()
}

type PostRouter struct {
}

func NewPostRouter() PostRouterHdl {
	return &PostRouter{}
}
