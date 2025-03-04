package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
)

type NumberRouterHdl interface {
	RegisterNumberRouters()
}

type NumberRouter struct {
	nc *controller.NumberController
	j  *middleware.Jwt
	e  *gin.Engine
}

func NewNumberRouter(nc *controller.NumberController, e *gin.Engine, j *middleware.Jwt) *NumberRouter {
	return &NumberRouter{
		nc: nc,
		e:  e,
		j:  j,
	}
}

func (nr *NumberRouter) RegisterNumberRouters() {
	number := nr.e.Group("/number")
	number.Use(nr.j.WrapCheckToken())
	{
		number.POST("/create", nr.nc.Send())
		number.POST("/delete", nr.nc.Delete())
		number.POST("/search", nr.nc.Search())
	}
}
