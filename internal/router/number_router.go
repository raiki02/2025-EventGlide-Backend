package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
)

type NumberRouterHdl interface {
	RegisterNumberRouters()
}

type NumberRouter struct {
	nc *controller.NumberController
	e  *gin.Engine
}

func NewNumberRouter(nc *controller.NumberController, e *gin.Engine) *NumberRouter {
	return &NumberRouter{
		nc: nc,
		e:  e,
	}
}

func (nr *NumberRouter) RegisterNumberRouters() {
	number := nr.e.Group("/number")
	{
		number.POST("/likes", nr.nc.SendLikesNum())
		number.POST("/comments", nr.nc.SendCommentsNum())
	}
}
