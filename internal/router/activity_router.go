package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
)

type ActRouterHdl interface {
	RegisterActRouters() error
}

type ActRouter struct {
	e   *gin.Engine
	ach *controller.ActController
}

func NewActRouter(e *gin.Engine, ach *controller.ActController) *ActRouter {
	return &ActRouter{
		e:   e,
		ach: ach,
	}
}

func (ar ActRouter) RegisterActRouters() error {
	act := ar.e.Group("act")
	{
		act.POST("/new", ar.ach.NewAct())
		act.POST("/draft", ar.ach.NewDraft())

		act.GET("/name", ar.ach.FindActByName())
		act.POST("/search", ar.ach.FindActBySearches())
		act.POST("/details", ar.ach.ShowActDetails())
	}
	return nil
}
