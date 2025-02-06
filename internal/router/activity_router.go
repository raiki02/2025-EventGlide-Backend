package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller/activity"
)

type ActRouterHdl interface {
	RegisterActRouter() error
}

type ActRouter struct {
	e   *gin.Engine
	ech activity.ActControllerHdl
}

func NewActRouter(e *gin.Engine, ech activity.ActControllerHdl) ActRouterHdl {
	return &ActRouter{
		e:   e,
		ech: ech,
	}
}

func (ar ActRouter) RegisterActRouter() error {
	act := ar.e.Group("act")
	{
		//1
		act.POST("/new", ar.ech.NewAct())
		act.POST("/draft", ar.ech.NewDraft())

		//0 or 1
		act.GET("/host", ar.ech.FindActByHost())
		act.GET("/type", ar.ech.FindActByType())
		act.GET("/location", ar.ech.FindActByLocation())
		act.GET("/signup", ar.ech.FindActByIfSignup())
		act.GET("/foreign", ar.ech.FindActByIsForeign())

		//more complex
		act.GET("/time", ar.ech.FindActByTime())
		act.GET("/name", ar.ech.FindActByName())
		act.GET("/:date", ar.ech.FindActByDate())
	}
	return nil
}
