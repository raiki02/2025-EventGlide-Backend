package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
)

type ActRouterHdl interface {
	RegisterActRouters() error
}

type ActRouter struct {
	e   *gin.Engine
	ach *controller.ActController
	j   *middleware.Jwt
}

func NewActRouter(e *gin.Engine, ach *controller.ActController, j *middleware.Jwt) *ActRouter {
	return &ActRouter{
		e:   e,
		ach: ach,
		j:   j,
	}
}

func (ar ActRouter) RegisterActRouters() error {
	act := ar.e.Group("act")
	act.Use(ar.j.WrapCheckToken())
	{
		act.POST("/create", ar.ach.NewAct())
		act.POST("/draft", ar.ach.NewDraft())
		act.POST("/load", ar.ach.LoadDraft())
		act.GET("/name/:name", ar.ach.FindActByName())
		act.GET("/date/:date", ar.ach.FindActByDate())
		act.POST("/search", ar.ach.FindActBySearches())
		act.GET("/own/:id", ar.ach.FindActByOwnerID())
		act.GET("/all", ar.ach.ListAllActs())
	}
	return nil
}
