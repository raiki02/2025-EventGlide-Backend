package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
)

type PostRouterHdl interface {
	RegisterPostRouters()
}

type PostRouter struct {
	e   *gin.Engine
	pch controller.PostControllerHdl
}

func NewPostRouter(e *gin.Engine, pch controller.PostControllerHdl) PostRouterHdl {
	return &PostRouter{
		e:   e,
		pch: pch,
	}
}

func (pr *PostRouter) RegisterPostRouters() {
	post := pr.e.Group("/post")
	{
		post.GET("/list", pr.pch.GetAllPost())
		post.POST("/create", pr.pch.CreatePost())
		post.GET("/find", pr.pch.FindPostByName())
		post.DELETE("/delete", pr.pch.DeletePost())
	}
}
