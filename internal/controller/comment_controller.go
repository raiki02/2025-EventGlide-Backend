package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
)

type CommentControllerHdl interface {
	CreateComment() gin.HandlerFunc
}

type CommentController struct {
	cd dao.CommentDAOHdl
}

func NewCommentController(cd dao.CommentDAOHdl) *CommentController {
	return &CommentController{cd}
}

func (cc *CommentController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment model.Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			tools.ReturnMSG(c, "bind error", err.Error())
			return
		}
		if err := cc.cd.Create(c, comment); err != nil {
			tools.ReturnMSG(c, "create error", err.Error())
			return
		}
		c.JSON(200, gin.H{"message": "comment created"})
	}
}
