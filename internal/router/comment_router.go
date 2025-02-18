package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
)

type CommentRouterHdl interface {
	RegisterCommentRouter()
}

type CommentRouter struct {
	cch *controller.CommentController
	e   *gin.Engine
}

func NewCommentRouter(cch *controller.CommentController, e *gin.Engine) *CommentRouter {
	return &CommentRouter{
		cch: cch,
		e:   e,
	}
}

func (cr *CommentRouter) RegisterCommentRouter() {
	cmt := cr.e.Group("/comment")
	{
		cmt.POST("/create", cr.cch.CreateComment())
		cmt.POST("/delete", cr.cch.DeleteComment())
		cmt.POST("/answer", cr.cch.AnswerComment())
	}
}
