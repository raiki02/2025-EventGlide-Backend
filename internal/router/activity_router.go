package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
)

type ActivityRouterHdl interface {
	RegisterActivityRouters()
}

type ActivityRouter struct {
	e   *gin.Engine
	ach controller.ActivityControllerHdl
	ch  middleware.ClaimsHdl
}

func NewActivityRouter(e *gin.Engine, ach controller.ActivityControllerHdl, ch middleware.ClaimsHdl) ActivityRouter {
	return ActivityRouter{e: e, ach: ach, ch: ch}
}

func (a *ActivityRouter) RegisterActivityRouters() {
	act := a.e.Group("/activity")
	{
		act.POST("/new", a.ach.NewActivity())
		act.POST("/draft", a.ach.NewDraft())
		act.GET("/all", a.ach.ListAllActivity())
		act.GET("/type", a.ach.ListActivityByType())
		act.GET("/time", a.ach.ListActivityByTime())
		act.GET("/host", a.ach.ListActivityByHost())
		act.GET("/location", a.ach.ListActivityByLocation())
		act.GET("/name", a.ach.ListActivityByName())
	}
}
