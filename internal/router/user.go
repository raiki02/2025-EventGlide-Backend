package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/pkg/ginx"
)

type UserRouterHdl interface {
	RegisterUserRouters()
}

type UserRouter struct {
	e  *gin.Engine
	j  *middleware.Jwt
	uc *controller.UserController
}

func NewUserRouter(e *gin.Engine, uc *controller.UserController, j *middleware.Jwt) *UserRouter {
	return &UserRouter{e: e, uc: uc, j: j}
}

func (ur *UserRouter) RegisterUserRouters() {
	user := ur.e.Group("/user")
	{
		user.POST("/login", ginx.WrapRequest(ur.uc.Login))

		user.Use(ur.j.WrapCheckToken())
		{
			user.POST("/logout", ginx.Wrap(ur.uc.Logout))
			user.GET("/token/qiniu", ginx.Wrap(ur.uc.GenQiniuToken))
			user.GET("/info/:id", ginx.WrapRequest(ur.uc.GetUserInfo))
			user.POST("/avatar", ginx.WrapRequestWithClaims(ur.uc.UpdateAvatar))
			user.POST("/username", ginx.WrapRequestWithClaims(ur.uc.UpdateUsername))
			user.POST("/search/act", ginx.WrapRequestWithClaims(ur.uc.SearchUserAct))
			user.POST("/search/post", ginx.WrapRequestWithClaims(ur.uc.SearchUserPost))
			user.POST("/collect/act", ginx.WrapWithClaims(ur.uc.LoadCollectAct))
			user.POST("/collect/post", ginx.WrapWithClaims(ur.uc.LoadCollectPost))
			user.POST("/like/act", ginx.WrapWithClaims(ur.uc.LoadLikeAct))
			user.POST("/like/post", ginx.WrapWithClaims(ur.uc.LoadLikePost))
			user.GET("/checking", ginx.WrapWithClaims(ur.uc.Checking))
		}
	}
}
