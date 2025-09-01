package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/middleware"
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
		user.POST("/login", ur.uc.Login())
		user.POST("/logout", ur.j.WrapCheckToken(), ur.uc.Logout())
		user.GET("/token/qiniu", ur.j.WrapCheckToken(), ur.uc.GenQiniuToken())
		user.GET("/info/:id", ur.j.WrapCheckToken(), ur.uc.GetUserInfo())
		user.POST("/avatar", ur.j.WrapCheckToken(), ur.uc.UpdateAvatar())
		user.POST("/username", ur.j.WrapCheckToken(), ur.uc.UpdateUsername())
		user.POST("/search/act", ur.j.WrapCheckToken(), ur.uc.SearchUserAct())
		user.POST("/search/post", ur.j.WrapCheckToken(), ur.uc.SearchUserPost())
		user.POST("/collect/act", ur.j.WrapCheckToken(), ur.uc.LoadCollectAct())
		user.POST("/collect/post", ur.j.WrapCheckToken(), ur.uc.LoadCollectPost())
		user.POST("/like/act", ur.j.WrapCheckToken(), ur.uc.LoadLikeAct())
		user.POST("/like/post", ur.j.WrapCheckToken(), ur.uc.LoadLikePost())
		user.GET("/checking", ur.j.WrapCheckToken(), ur.uc.Checking())
	}
}
