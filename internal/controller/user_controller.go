package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/api/resp"
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
	GenQINIUToken() gin.HandlerFunc
	Like() gin.HandlerFunc
	Comment() gin.HandlerFunc
	Collect() gin.HandlerFunc
	LoadCollect() gin.HandlerFunc
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
// @Param user body req.LoginReq true "登录请求"
// @Success 200 {object} resp.Resp{data=resp.LoginResp}
// @Router /user/login [post]
func (uc *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var lr req.LoginReq
		err := c.ShouldBindJSON(&lr)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if lr.StudentID == "" || lr.Password == "" {
			c.JSON(200, tools.ReturnMSG(c, "studentid or password is empty", nil))
			return
		}
		user, token, err := uc.ush.Login(c, lr.StudentID, lr.Password)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "login fail", nil))
			return
		}
		res := resp.LoginResp{
			Id:     user.Id,
			Sid:    user.StudentID,
			Name:   user.Name,
			Avatar: user.Avatar,
			School: user.School,
			Token:  token,
		}
		c.JSON(200, tools.ReturnMSG(c, "success", res))
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
		c.JSON(200, tools.ReturnMSG(c, "success", nil))
	}
}

// @Tags User
// @Summary 获取用户信息
// @Produce json
// @Param Authorization header string true "token"
// @Param id path string true "用户id"
// @Success 200 {object} resp.Resp{data=model.User}
// @Router /user/info/{id} [get]
func (uc *UserController) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Param("id")
		if sid == "" {
			c.JSON(200, tools.ReturnMSG(c, "id is empty", nil))
			return
		}
		user, err := uc.ush.GetUserInfo(c, sid)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "get user info fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", user))
	}
}

// @Tags User
// @Summary 更新头像
// @Description not finished
// @Produce json
// @Param Authorization header string true "token"
// @Param userAvatarReq body req.UserAvatarReq true "用户头像更改"
// @Success 200 {object} resp.Resp
// @Router /user/avatar [post]
func (uc *UserController) UpdateAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userAvatarReq req.UserAvatarReq
		err := c.ShouldBind(&userAvatarReq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if userAvatarReq.StudentID == "" || userAvatarReq.AvatarUrl == "" {
			c.JSON(200, tools.ReturnMSG(c, "param empty", nil))
			return
		}
		err = uc.ush.UpdateAvatar(c, userAvatarReq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "update avatar fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", nil))
	}
}

// @Tags User
// @Summary 更新用户名
// @Produce json
// @Param Authorization header string true "token"
// @Param unr body req.UpdateNameReq true "更新用户名"
// @Success 200 {object} resp.Resp
// @Router /user/username [post]
func (uc *UserController) UpdateUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		var unr req.UpdateNameReq
		err := c.ShouldBind(&unr)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if unr.Name == "" {
			c.JSON(200, tools.ReturnMSG(c, "name is empty", nil))
			return
		}
		err = uc.ush.UpdateUsername(c, unr.StudentID, unr.Name)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "update username fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", nil))
	}

}

// @Tags User
// @Summary 搜索用户活动
// @Produce json
// @Param Authorization header string true "token"
// @Param ureq body req.UserSearchReq true "搜索请求"
// @Success 200 {object} resp.Resp{data=[]model.Activity}
// @Router /user/search/act [post]
func (uc *UserController) SearchUserAct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ureq req.UserSearchReq
		err := c.ShouldBindJSON(&ureq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if ureq.StudentID == "" {
			c.JSON(200, tools.ReturnMSG(c, "sid is empty", nil))
			return
		}
		acts, err := uc.ush.SearchUserAct(c, ureq.StudentID, ureq.Keyword)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "search user act fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", acts))
	}
}

// @Tags User
// @Summary 搜索用户帖子
// @Produce json
// @Param Authorization header string true "token"
// @Param ureq body req.UserSearchReq true "搜索请求"
// @Success 200 {object} resp.Resp{data=[]model.Post}
// @Router /user/search/post [post]
func (uc *UserController) SearchUserPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ureq req.UserSearchReq
		err := c.ShouldBindJSON(&ureq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if ureq.StudentID == "" {
			c.JSON(200, tools.ReturnMSG(c, "sid is empty", nil))
			return
		}
		posts, err := uc.ush.SearchUserPost(c, ureq.StudentID, ureq.Keyword)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "search user post fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", posts))
	}
}

// @Tags User
// @Summary 获取七牛云token
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp{data=resp.ImgBedResp}
// @Router /user/token/qiniu [get]
func (uc *UserController) GenQiniuToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := uc.ush.GenQINIUToken(c)
		c.JSON(200, tools.ReturnMSG(c, "success", res))
	}
}

// @Tags User
// @Summary 点赞
// @Produce json
// @Param Authorization header string true "token"
// @Param lr body req.NumReq true "点赞请求"
// @Success 200 {object} resp.Resp
// @Router /user/like [post]
func (uc *UserController) Like() gin.HandlerFunc {
	return func(c *gin.Context) {
		var lr req.NumReq
		err := c.ShouldBindJSON(&lr)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		err = uc.ush.Like(c, lr)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "like fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", nil))
	}
}

// @Tags User
// @Summary 评论
// @Produce json
// @Param Authorization header string true "token"
// @Param cr body req.NumReq true "评论请求"
// @Success 200 {object} resp.Resp
// @Router /user/comment [post]
func (uc *UserController) Comment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cr req.NumReq
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		err = uc.ush.Comment(c, cr.TargetId, cr.Object, cr.Type)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "comment fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", nil))
	}
}

// @Tags User
// @Summary 收藏
// @Produce json
// @Param Authorization header string true "token"
// @Param cr body req.NumReq true "收藏请求"
// @Success 200 {object} resp.Resp
// @Router /user/collect [post]
func (uc *UserController) Collect() gin.HandlerFunc {
	return func(context *gin.Context) {
		var cr req.NumReq
		err := context.ShouldBindJSON(&cr)
		if err != nil {
			context.JSON(200, tools.ReturnMSG(context, err.Error(), nil))
			return
		}
		err = uc.ush.Collect(context, cr)
		if err != nil {
			context.JSON(200, tools.ReturnMSG(context, "collect fail", nil))
			return
		}
		context.JSON(200, tools.ReturnMSG(context, "success", nil))
	}
}

// @Tags User
// @Summary 加载活动收藏
// @Produce json
// @Param Authorization header string true "token"
// @Param cr body req.NumReq true "加载收藏请求"
// @Success 200 {object} resp.Resp{data=[]resp.ListActivitiesResp}
// @Router /user/collect/act [post]
func (uc *UserController) LoadCollectAct() gin.HandlerFunc {
	return func(context *gin.Context) {
		var cr req.NumReq
		err := context.ShouldBindJSON(&cr)
		if err != nil {
			context.JSON(200, tools.ReturnMSG(context, err.Error(), nil))
			return
		}
		if cr.StudentID == "" {
			context.JSON(200, tools.ReturnMSG(context, "studentid is empty", nil))
			return
		}
		res, err := uc.ush.LoadCollectAct(context, cr.StudentID)
		if err != nil {
			context.JSON(200, tools.ReturnMSG(context, "load collect fail", nil))
			return
		}
		context.JSON(200, tools.ReturnMSG(context, "success", res))
	}
}

// @Tags User
// @Summary 加载帖子收藏
// @Produce json
// @Param Authorization header string true "token"
// @Param cr body req.NumReq true "加载收藏请求"
// @Success 200 {object} resp.Resp{data=[]resp.ListPostsResp}
// @Router /user/collect/post [post]
func (uc *UserController) LoadCollectPost() gin.HandlerFunc {
	return func(context *gin.Context) {
		var cr req.NumReq
		err := context.ShouldBindJSON(&cr)
		if err != nil {
			context.JSON(200, tools.ReturnMSG(context, err.Error(), nil))
			return
		}
		if cr.StudentID == "" {
			context.JSON(200, tools.ReturnMSG(context, "studentid is empty", nil))
			return
		}
		res, err := uc.ush.LoadCollectPost(context, cr.StudentID)
		if err != nil {
			context.JSON(200, tools.ReturnMSG(context, "load collect fail", nil))
			return
		}
		context.JSON(200, tools.ReturnMSG(context, "success", res))
	}
}
