package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/raiki02/EG/docs"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/internal/service"
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
	fr   *FeedRouter
	ir   *InteractionRouter
	cors *middleware.Cors

	kgr ApiKeyRouter
	cba service.CallbackAuditorService
}

func NewRouter(e *gin.Engine, ur *UserRouter, ar *ActRouter, pr *PostRouter, cr *CommentRouter, fr *FeedRouter, ir *InteractionRouter, cors *middleware.Cors, cba service.CallbackAuditorService, kgr ApiKeyRouter) *Router {
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
		kgr:  kgr,
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
	r.cba.RegisterCallbackAuditorRouters()
	r.kgr.RegisterApiKeyRouters()
	r.RegisterSwagger()
}

func (r *Router) Run() error {
	return r.e.Run()
}

func (r *Router) RegisterSwagger() {
	r.e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
