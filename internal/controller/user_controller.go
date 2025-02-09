package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/service"
)

// user这边操作数据库不频繁
type UserControllerHdl interface {
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	GetUserInfo() gin.HandlerFunc
	UpdateAvatar() gin.HandlerFunc
	UpdateUsername() gin.HandlerFunc
	SearchUserAct() gin.HandlerFunc
	SearchUserPost() gin.HandlerFunc
}

type UserController struct {
	e   *gin.Engine
	ush service.UserServiceHdl
}

func NewUserController(e *gin.Engine, ush service.UserServiceHdl) UserControllerHdl {
	return &UserController{
		e:   e,
		ush: ush,
	}
}

func (uc *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (uc *UserController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (uc *UserController) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (uc *UserController) UpdateAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (uc *UserController) UpdateUsername() gin.HandlerFunc {
	return func(c *gin.Context) {

	}

}

func (uc *UserController) SearchUserAct() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (uc *UserController) SearchUserPost() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
