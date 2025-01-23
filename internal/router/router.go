package router

import "github.com/gin-gonic/gin"

type RouterHdl interface {
	RegisterRouters()
}

type Router struct {
	e  *gin.Engine
	ur UserRouterHdl
	ar ActivityRouterHdl
	pr PostRouterHdl
	cr CommentRouterHdl
}

func NewRouter(e *gin.Engine, ur UserRouterHdl, ar ActivityRouterHdl, pr PostRouterHdl, cr CommentRouterHdl) Router {
	return Router{
		e:  e,
		ur: ur,
		ar: ar,
		pr: pr,
		cr: cr}
}

func (r *Router) RegisterRouters() {
	r.ur.RegisterUserRouters()
	r.ar.RegisterActivityRouters()
	r.pr.RegisterPostRouters()
	r.cr.RegisterCommentRouter()
}
