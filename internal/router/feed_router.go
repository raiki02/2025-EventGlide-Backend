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
