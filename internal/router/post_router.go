package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
)

type PostRouterHdl interface {
	RegisterPostRouters()
}

type PostRouter struct {
	e   *gin.Engine
	j   *middleware.Jwt
	pch *controller.PostController
}

func NewPostRouter(e *gin.Engine, pch *controller.PostController, j *middleware.Jwt) *PostRouter {
	return &PostRouter{
		e:   e,
		pch: pch,
		j:   j,
	}
}

func (pr *PostRouter) RegisterPostRouters() {
	post := pr.e.Group("/post")
	post.Use(pr.j.WrapCheckToken())
	{
		post.GET("/list", pr.pch.GetAllPost())
		post.POST("/create", pr.pch.CreatePost())
		post.GET("/find", pr.pch.FindPostByName())
		post.POST("/draft", pr.pch.CreateDraft())
		post.DELETE("/delete", pr.pch.DeletePost())
		post.POST("/load", pr.pch.LoadDraft())
	}
}
