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
// @Param CommentReq body req.CommentReq true "评论"
// @Success 200 {object} resp.Resp{data=resp.CommentResp}
// @Router /comment/create [post]
func (cc *CommentController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cReq req.CommentReq
		err := c.ShouldBindJSON(&cReq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}
		cmt, err := cc.cs.CreateComment(c, &cReq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "create comment fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "create comment success", cmt))
	}
}

// @Tags Comment
// @Summary 回复评论
// @Produce json
// @Param Authorization header string true "token"
// @Param CommentReq body req.CommentReq true "回复"
// @Success 200 {object} resp.Resp{data=resp.AnswerResp}
// @Router /comment/answer [post]
func (cc *CommentController) AnswerComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cReq req.CommentReq
		err := c.ShouldBindJSON(&cReq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}
		asw, err := cc.cs.AnswerComment(c, &cReq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "answer comment fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "answer comment success", asw))
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
		var dcr req.DeleteCommentReq
		err := c.ShouldBindJSON(&dcr)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}

		if dcr.TargetID == "" || dcr.Sid == "" {
			c.JSON(200, tools.ReturnMSG(c, "param error", nil))
			return
		}
		err = cc.cs.DeleteComment(c, dcr.Sid, dcr.TargetID)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "delete comment fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "delete comment success", nil))
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
		target := c.Param("id")
		if target == "" {
			c.JSON(200, tools.ReturnMSG(c, "param empty", nil))
			return
		}
		comments, err := cc.cs.LoadComments(c, target)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "load comments fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "load comments success", comments))
	}
}

// @Tags Comment
// @Summary 加载回复
// @Produce json
// @Param id path string true "目标id"
// @Success 200 {object} resp.Resp{data=[]resp.AnswerResp}
// @Router /comment/answer/{id} [get]
func (cc *CommentController) LoadAnswers() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := c.Param("id")
		if target == "" {
			c.JSON(200, tools.ReturnMSG(c, "param empty", nil))
			return
		}
		answers, err := cc.cs.LoadAnswers(c, target)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, "load answers fail", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "load answers success", answers))
	}
}
