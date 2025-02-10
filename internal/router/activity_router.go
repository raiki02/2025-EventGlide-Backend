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
		//1
		act.POST("/new", ar.ach.NewAct())
		act.POST("/draft", ar.ach.NewDraft())

		//0 or 1
		act.GET("/host", ar.ach.FindActByHost())
		act.GET("/type", ar.ach.FindActByType())
		act.GET("/location", ar.ach.FindActByLocation())
		act.GET("/signup", ar.ach.FindActByIfSignup())
		act.GET("/foreign", ar.ach.FindActByIsForeign())

		//more complex
		act.GET("/time", ar.ach.FindActByTime())
		act.GET("/name", ar.ach.FindActByName())
		act.GET("/:date", ar.ach.FindActByDate())
	}
	return nil
}
