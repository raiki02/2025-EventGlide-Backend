package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

type CommentControllerHdl interface {
	CreateComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc
	AnswerComment() gin.HandlerFunc
	LoadComments() gin.HandlerFunc
	LoadAnswers() gin.HandlerFunc
}

type CommentController struct {
	cs *service.CommentService
}

func NewCommentController(cs *service.CommentService) *CommentController {
	return &CommentController{
		cs: cs,
	}
}

// @Tags Comment
// @Summary 创建评论
// @Produce json
// @Param Authorization header string true "token"
// @Param CommentReq body req.CreateCommentReq true "评论"
// @Success 200 {object} resp.Resp{data=resp.CommentResp}
// @Router /comment/create [post]
func (cc *CommentController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r req.CreateCommentReq
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if r.StudentID == "" || r.Content == "" || r.ParentID == "" {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}
		res, err := cc.cs.CreateComment(c, r)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", res))
	}
}

// @Tags Comment
// @Summary 回复评论
// @Produce json
// @Param Authorization header string true "token"
// @Param CommentReq body req.CreateCommentReq true "回复"
// @Success 200 {object} resp.Resp{data=resp.CommentResp}
// @Router /comment/answer [post]
func (cc *CommentController) AnswerComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r req.CreateCommentReq
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if r.StudentID == "" || r.Content == "" || r.ParentID == "" {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}
		res, err := cc.cs.AnswerComment(c, r)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", res))
	}
}

// @Tags Comment
// @Summary 删除评论
// @Produce json
// @Param sid formData string true "学号"
// @Param Authorization header string true "token"
// @Param DeleteCommentReq body req.DeleteCommentReq true "删除评论"
// @Success 200 {object} resp.Resp
// @Router /comment/delete [post]
func (cc *CommentController) DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r req.DeleteCommentReq
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if r.StudentID == "" || r.TargetID == "" {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}
		err = cc.cs.DeleteComment(c, r)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", nil))
	}
}

// @Tags Comment
// @Summary 加载评论
// @Produce json
// @Param id path string true "目标id"
// @Success 200 {object} resp.Resp{data=[]resp.CommentResp}
// @Router /comment/load/{id} [get]
func (cc *CommentController) LoadComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}
		res, err := cc.cs.LoadComments(c, id)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", res))

	}
}

// @Tags Comment
// @Summary 加载回复
// @Produce json
// @Param id path string true "目标id"
// @Success 200 {object} resp.Resp{data=[]resp.CommentResp}
// @Router /comment/answer/{id} [get]
func (cc *CommentController) LoadAnswers() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}
		res, err := cc.cs.LoadAnswers(c, id)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", res))
	}
}
