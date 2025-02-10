package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
)

type CommentServiceHdl interface {
	CreateComment(*gin.Context, *model.Comment) error
	DeleteComment(*gin.Context, string) error
	AnswerComment(*gin.Context, *model.Comment) error
}

type CommentService struct {
	cd *dao.CommentDao
}

func NewCommentService(cd *dao.CommentDao) *CommentService {
	return &CommentService{
		cd: cd,
	}
}

func (cs *CommentService) CreateComment(c *gin.Context, cmt *model.Comment) error {
	cmt.CommentID = tools.GenUUID(c)
	return cs.cd.CreateComment(c, cmt)
}

func (cs *CommentService) DeleteComment(c *gin.Context, commentID string) error {
	return cs.cd.DeleteComment(c, commentID)
}

func (cs *CommentService) AnswerComment(c *gin.Context, cmt *model.Comment) error {
	cmt.CommentID = tools.GenUUID(c)
	return cs.cd.AnswerComment(c, cmt)
}
