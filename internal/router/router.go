package router

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/raiki02/EG/docs"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
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
	fr   *FeedRouter
	ir   *InteractionRouter
	cors *middleware.Cors

	cba *controller.CallbackAuditorController
}

func NewRouter(e *gin.Engine, ur *UserRouter, ar *ActRouter, pr *PostRouter, cr *CommentRouter, fr *FeedRouter, ir *InteractionRouter, cors *middleware.Cors, cba *controller.CallbackAuditorController) *Router {
	return &Router{
		e:    e,
		ur:   ur,
		ar:   ar,
		pr:   pr,
		cr:   cr,
		fr:   fr,
		ir:   ir,
		cors: cors,
		cba:  cba,
	}
}

func (r *Router) RegisterRouters() {
	r.cors.HandleCors()
	r.ur.RegisterUserRouters()
	r.ar.RegisterActRouters()
	r.pr.RegisterPostRouters()
	r.cr.RegisterCommentRouter()
	r.ir.RegisterInteractionRouters()
	r.fr.RegisterFeedRouters()
	r.RegisterSwagger()
}

func (r *Router) Run() (error, func()) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.e.Handler(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return nil, func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}
}

func (r *Router) RegisterSwagger() {
	r.e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
