package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

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

// @Tags User
// @Summary 登录
// @Produce json
// @Param studentid formData string true "学号"
// @Param password formData string true "密码"
// @Success 200 {object} resp.Resp
// @Router /user/login [post]
func (uc *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentid := c.PostForm("studentid")
		password := c.PostForm("password")
		if studentid == "" || password == "" {
			c.JSON(200, tools.ReturnMSG(c, "studentid or password is empty", nil))
			return
		}
		user, token, err := uc.ush.Login(c, studentid, password)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "login fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "login success", user, token))
	}
}

// @Tags User
// @Summary 登出
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp
// @Router /user/logout [post]
func (uc *UserController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		err := uc.ush.Logout(c, token)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "logout fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "logout success", nil))
	}
}

// @Tags User
// @Summary 获取用户信息
// @Produce json
// @Param sid query string true "学号"
// @Success 200 {object} resp.Resp
// @Router /user/info [get]
func (uc *UserController) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Query("sid")
		if sid == "" {
			c.JSON(200, tools.ReturnMSG(c, "sid is empty", nil))
			return
		}
		user, err := uc.ush.GetUserInfo(c, sid)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "get user info fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "get user info success", user))
	}
}

// @Tags User
// @Summary 更新头像
// @Description not finished
// @Produce json
// @Param sid formData string true "学号"
// @Param file formData file true "图片"
// @Success 200 {object} resp.Resp
// @Router /user/avatar [post]
func (uc *UserController) UpdateAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.PostForm("sid")
		err := uc.ush.UpdateAvatar(c, sid)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "update avatar fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "update avatar success", nil))
	}
}

// @Tags User
// @Summary 更新用户名
// @Produce json
// @Param sid formData string true "学号"
// @Param newname formData string true "新用户名"
// @Success 200 {object} resp.Resp
// @Router /user/username [post]
func (uc *UserController) UpdateUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.PostForm("sid")
		name := c.PostForm("newname")
		if name == "" {
			c.JSON(200, tools.ReturnMSG(c, "name is empty", nil))
			return
		}
		err := uc.ush.UpdateUsername(c, sid, name)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "update username fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "update username success", nil))
	}

}

// @Tags User
// @Summary 搜索用户活动
// @Produce json
// @Param sid query string true "学号"
// @Param keyword query string true "关键字"
// @Success 200 {object} resp.Resp
// @Router /user/search/act [get]
func (uc *UserController) SearchUserAct() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Query("sid")
		keyword := c.Query("keyword")
		if sid == "" {
			c.JSON(200, tools.ReturnMSG(c, "sid is empty", nil))
			return
		}
		acts, err := uc.ush.SearchUserAct(c, sid, keyword)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "search user act fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "search user act success", acts))
	}
}

// @Tags User
// @Summary 搜索用户帖子
// @Produce json
// @Param sid query string true "学号"
// @Param keyword query string true "关键字"
// @Success 200 {object} resp.Resp
// @Router /user/search/post [get]
func (uc *UserController) SearchUserPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Query("sid")
		keyword := c.Query("keyword")
		if sid == "" {
			c.JSON(200, tools.ReturnMSG(c, "sid is empty", nil))
			return
		}
		posts, err := uc.ush.SearchUserPost(c, sid, keyword)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "search user post fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "search user post success", posts))
	}
}
