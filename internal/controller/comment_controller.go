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

func (cc *CommentController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmt model.Comment
		//content, posterID, TargetID
		if err := c.ShouldBindJSON(&cmt); err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		if err := cc.cs.CreateComment(c, &cmt); err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "comment success", nil)
	}
}

func (cc *CommentController) DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.PostForm("comment_id")
		if cid == "" {
			tools.ReturnMSG(c, "comment not exist", nil)
			return
		}
		if err := cc.cs.DeleteComment(c, cid); err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "delete success", nil)
	}
}

func (cc *CommentController) AnswerComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmt model.Comment
		//content, posterID, TargetID, ParentID
		if err := c.ShouldBindJSON(&cmt); err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		if err := cc.cs.AnswerComment(c, &cmt); err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "answer success", nil)
	}
}
