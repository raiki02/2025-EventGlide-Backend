package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
)

type CommentRouterHdl interface {
	RegisterCommentRouter()
}

type CommentRouter struct {
	cch *controller.CommentController
	e   *gin.Engine
	j   *middleware.Jwt
}

func NewCommentRouter(cch *controller.CommentController, e *gin.Engine, j *middleware.Jwt) *CommentRouter {
	return &CommentRouter{
		cch: cch,
		e:   e,
		j:   j,
	}
}

func (cr *CommentRouter) RegisterCommentRouter() {
	cmt := cr.e.Group("/comment")
	cmt.Use(cr.j.WrapCheckToken())
	{
		cmt.POST("/create", cr.cch.CreateComment())
		cmt.POST("/delete", cr.cch.DeleteComment())
		cmt.POST("/answer", cr.cch.AnswerComment())
	}
}
