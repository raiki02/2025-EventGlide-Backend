package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
)

type UserRouterHdl interface {
	RegisterUserRouters()
}

type UserRouter struct {
	e  *gin.Engine
	uc *controller.UserController
}

func NewUserRouter(e *gin.Engine, uc *controller.UserController) *UserRouter {
	return &UserRouter{e: e, uc: uc}
}

func (ur *UserRouter) RegisterUserRouters() {
	user := ur.e.Group("/user")
	{
		user.POST("/login", ur.uc.Login())
		user.POST("/logout", ur.uc.Logout())
		user.GET("/info", ur.uc.GetUserInfo())
		user.POST("/avatar", ur.uc.UpdateAvatar())
		user.POST("/username", ur.uc.UpdateUsername())
		user.POST("/search/act", ur.uc.SearchUserAct())
		user.POST("/search/post", ur.uc.SearchUserPost())
	}
}
