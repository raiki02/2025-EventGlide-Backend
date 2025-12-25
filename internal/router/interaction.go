package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/pkg/ginx"
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
		i.POST("/like", ginx.WrapRequestWithClaims(ir.ic.Like))
		i.POST("/dislike", ginx.WrapRequestWithClaims(ir.ic.Dislike))

		i.POST("/collect", ginx.WrapRequestWithClaims(ir.ic.Collect))
		i.POST("/discollect", ginx.WrapRequestWithClaims(ir.ic.Discollect))

		i.POST("/approve", ginx.WrapRequestWithClaims(ir.ic.Approve))
		i.POST("/reject", ginx.WrapRequestWithClaims(ir.ic.Reject))
	}
}
