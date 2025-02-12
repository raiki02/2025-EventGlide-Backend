package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

type CommentControllerHdl interface {
	CreateComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc
	AnswerComment() gin.HandlerFunc
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
// @Accept json
// @Param comment body model.Comment true "评论"
// @Success 200 {object} resp.Resp
// @Router /comment/create [post]
func (cc *CommentController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmt model.Comment
		//content, posterID, TargetID
		if err := c.ShouldBindJSON(&cmt); err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if err := cc.cs.CreateComment(c, &cmt); err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "comment success", nil))
	}
}

// @Tags Comment
// @Summary 删除评论
// @Produce json
// @Param comment_id formData string true "评论ID"
// @Success 200 {object} resp.Resp
// @Router /comment/delete [post]
func (cc *CommentController) DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.PostForm("comment_id")
		if cid == "" {
			c.JSON(200, tools.ReturnMSG(c, "comment not exist", nil))
			return
		}
		if err := cc.cs.DeleteComment(c, cid); err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "delete success", nil))
	}
}

// @Tags Comment
// @Summary 回复评论
// @Produce json
// @Accept json
// @Param comment body model.Comment true "评论"
// @Success 200 {object} resp.Resp
// @Router /comment/answer [post]
func (cc *CommentController) AnswerComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmt model.Comment
		//content, posterID, TargetID, ParentID
		if err := c.ShouldBindJSON(&cmt); err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if err := cc.cs.AnswerComment(c, &cmt); err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "answer success", nil))
	}
}
