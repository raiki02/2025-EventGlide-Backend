package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
)

type FeedRouterHdl interface {
	RegisterFeedRouters()
}

type FeedRouter struct {
	fc *controller.FeedController
	j  *middleware.Jwt
	e  *gin.Engine
}

func NewFeedRouter(fc *controller.FeedController, e *gin.Engine, j *middleware.Jwt) *FeedRouter {
	return &FeedRouter{
		fc: fc,
		e:  e,
		j:  j,
	}
}

func (fr *FeedRouter) RegisterFeedRouters() {
	feed := fr.e.Group("/feed")
	feed.Use(fr.j.WrapCheckToken())
	{
		feed.GET("/total", fr.fc.GetTotalCnt())
		feed.GET("/list", fr.fc.GetFeedList())
		feed.GET("/auditor", fr.fc.GetAuditorFeedList())
	}
}
