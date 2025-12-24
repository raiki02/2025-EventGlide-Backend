package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/pkg/ginx"
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
		feed.GET("/total", ginx.WrapWithClaims(fr.fc.GetTotalCnt))
		feed.GET("/list", ginx.WrapWithClaims(fr.fc.GetFeedList))
		feed.GET("/read/detail", ginx.WrapRequestWithClaims(fr.fc.ReadFeedDetail))
		feed.GET("/read/all", ginx.WrapWithClaims(fr.fc.ReadAllFeed))
		feed.GET("/auditor", ginx.WrapWithClaims(fr.fc.GetAuditorFeedList))
	}
}
