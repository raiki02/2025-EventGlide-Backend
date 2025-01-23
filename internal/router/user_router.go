package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/controller"
)

type UserRouterHdl interface {
	RegisterUserRouters()
}

type UserRouter struct {
	e  *gin.Engine
	uc controller.UserControllerHdl
}

func NewUserRouter(e *gin.Engine, uc controller.UserControllerHdl) UserRouterHdl {
	return &UserRouter{e: e, uc: uc}
}

func (ur *UserRouter) RegisterUserRouters() {
	ctx := context.Background()
	user := ur.e.Group("/user")
	{
		user.POST("/login", ur.uc.Login(ctx))
		user.POST("/logout", ur.uc.CheckToken(ctx), ur.uc.Logout(ctx))
		user.GET("/info", ur.uc.CheckToken(ctx), ur.uc.GetUserInfo(ctx))
		user.POST("/avatar", ur.uc.CheckToken(ctx), ur.uc.UpdateAvatar(ctx))
		user.POST("/username", ur.uc.CheckToken(ctx), ur.uc.UpdateUsername(ctx))
	}
}
