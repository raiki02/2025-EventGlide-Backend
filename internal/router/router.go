package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/middleware"
)

type RouterHdl interface {
	RegisterRouters()
	Run() error
}

type Router struct {
	e    *gin.Engine
	ur   *UserRouter
	ar   *ActRouter
	pr   *PostRouter
	cr   *CommentRouter
	cors *middleware.Cors
}

func NewRouter(e *gin.Engine, ur *UserRouter, ar *ActRouter, pr *PostRouter, cr *CommentRouter, cors *middleware.Cors) *Router {
	return &Router{
		e:    e,
		ur:   ur,
		ar:   ar,
		pr:   pr,
		cr:   cr,
		cors: cors,
	}
}

func (r *Router) RegisterRouters() {
	r.ur.RegisterUserRouters()
	r.ar.RegisterActRouters()
	r.pr.RegisterPostRouters()
	r.cr.RegisterCommentRouter()
}

func (r *Router) Run() error {
	return r.e.Run()
}
