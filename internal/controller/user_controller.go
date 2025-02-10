package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
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
	ush *service.UserService
}

func NewUserController(e *gin.Engine, ush *service.UserService) *UserController {
	return &UserController{
		e:   e,
		ush: ush,
	}
}

func (uc *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentid := c.PostForm("studentid")
		password := c.PostForm("password")
		if studentid == "" || password == "" {
			tools.ReturnMSG(c, "studentid or password is empty", nil)
			return
		}
		user, token, err := uc.ush.Login(c, studentid, password)
		if err != nil {
			tools.ReturnMSG(c, "login fail", nil)
			return
		}
		tools.ReturnMSG(c, "login success", user, token)
	}
}

func (uc *UserController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := uc.ush.Logout(c)
		if err != nil {
			tools.ReturnMSG(c, "logout fail", nil)
			return
		}
		tools.ReturnMSG(c, "logout success", nil)
	}
}

func (uc *UserController) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Query("sid")
		if sid == "" {
			tools.ReturnMSG(c, "sid is empty", nil)
			return
		}
		user, err := uc.ush.GetUserInfo(c, sid)
		if err != nil {
			tools.ReturnMSG(c, "get user info fail", nil)
			return
		}
		tools.ReturnMSG(c, "get user info success", user)
	}
}

func (uc *UserController) UpdateAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.PostForm("sid")
		err := uc.ush.UpdateAvatar(c, sid)
		if err != nil {
			tools.ReturnMSG(c, "update avatar fail", nil)
			return
		}
		tools.ReturnMSG(c, "update avatar success", nil)
	}
}

func (uc *UserController) UpdateUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.PostForm("sid")
		name := c.PostForm("newname")
		if name == "" {
			tools.ReturnMSG(c, "name is empty", nil)
			return
		}
		err := uc.ush.UpdateUsername(c, sid, name)
		if err != nil {
			tools.ReturnMSG(c, "update username fail", nil)
			return
		}
		tools.ReturnMSG(c, "update username success", nil)
	}

}

func (uc *UserController) SearchUserAct() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Query("sid")
		keyword := c.Query("keyword")
		if sid == "" {
			tools.ReturnMSG(c, "sid is empty", nil)
			return
		}
		acts, err := uc.ush.SearchUserAct(c, sid, keyword)
		if err != nil {
			tools.ReturnMSG(c, "search user act fail", nil)
			return
		}
		tools.ReturnMSG(c, "search user act success", acts)
	}
}

func (uc *UserController) SearchUserPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Query("sid")
		keyword := c.Query("keyword")
		if sid == "" {
			tools.ReturnMSG(c, "sid is empty", nil)
			return
		}
		posts, err := uc.ush.SearchUserPost(c, sid, keyword)
		if err != nil {
			tools.ReturnMSG(c, "search user post fail", nil)
			return
		}
		tools.ReturnMSG(c, "search user post success", posts)
	}
}
