package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/raiki02/EG/docs"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
	nr   *NumberRouter
	cors *middleware.Cors
}

func NewRouter(e *gin.Engine, ur *UserRouter, ar *ActRouter, pr *PostRouter, cr *CommentRouter, nr *NumberRouter, cors *middleware.Cors) *Router {
	return &Router{
		e:    e,
		ur:   ur,
		ar:   ar,
		pr:   pr,
		cr:   cr,
		nr:   nr,
		cors: cors,
	}
}

func (r *Router) RegisterRouters() {
	r.ur.RegisterUserRouters()
	r.ar.RegisterActRouters()
	r.pr.RegisterPostRouters()
	r.cr.RegisterCommentRouter()
	r.nr.RegisterNumberRouters()
	r.RegisterSwagger()
}

func (r *Router) Run() error {
	r.cors.HandleCors()
	r.RegisterRouters()
	return r.e.Run()
}

func (r *Router) RegisterSwagger() {
	r.e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
