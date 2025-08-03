package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
)

type InteractionRouter struct {
	e  *gin.Engine
	ic *controller.InteractionController
	j  *middleware.Jwt
}

func NewInteractionRouter(e *gin.Engine, ic *controller.InteractionController, j *middleware.Jwt) *InteractionRouter {
	return &InteractionRouter{
		e:  e,
		ic: ic,
		j:  j,
	}
}

func (ir *InteractionRouter) RegisterInteractionRouters() {
	i := ir.e.Group("interaction")
	i.Use(ir.j.WrapCheckToken())
	{
		i.POST("/like", ir.ic.Like())
		i.POST("/dislike", ir.ic.Dislike())

		i.POST("/collect", ir.ic.Collect())
		i.POST("/discollect", ir.ic.Discollect())
	}
}
