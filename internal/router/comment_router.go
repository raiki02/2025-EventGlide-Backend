package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
)

type CommentRouterHdl interface {
	RegisterCommentRouter()
}

type CommentRouter struct {
	cch controller.CommentControllerHdl
	e   *gin.Engine
}

func NewCommentRouter(cch controller.CommentControllerHdl, e *gin.Engine) *CommentRouter {
	return &CommentRouter{cch, e}
}

func (cr *CommentRouter) RegisterCommentRouter() {
	cmt := cr.e.Group("/comment")
	{
		cmt.POST("/create", cr.cch.CreateComment())
	}
}
