package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
	"time"
)

type CommentServiceHdl interface {
	CreateComment(*gin.Context, *req.CommentReq) error
	DeleteComment(*gin.Context, string, string) error
	AnswerComment(*gin.Context, *req.CommentReq) error
	LoadComments(*gin.Context, string) ([]model.Comment, error)
	LoadAnswers(*gin.Context, string) ([]model.SubComment, error)
}

type CommentService struct {
	cd *dao.CommentDao
}

func NewCommentService(cd *dao.CommentDao) *CommentService {
	return &CommentService{
		cd: cd,
	}
}

func (cs *CommentService) CreateComment(c *gin.Context, req *req.CommentReq) (*model.Comment, error) {
	cmt := toComment(c, req)
	cmt.Bid = tools.GenUUID(c)
	return cmt, cs.cd.CreateComment(c, cmt)
}

func (cs *CommentService) AnswerComment(c *gin.Context, req *req.CommentReq) (*model.SubComment, error) {
	answer := toAnswer(c, req)
	answer.Bid = tools.GenUUID(c)
	return answer, cs.cd.AnswerComment(c, answer)
}

func (cs *CommentService) DeleteComment(c *gin.Context, sid string, targetId string) error {
	return cs.cd.DeleteComment(c, sid, targetId)
}

func (cs *CommentService) LoadComments(c *gin.Context, targetId string) ([]model.Comment, error) {
	return cs.cd.LoadComments(c, targetId)
}

func (cs *CommentService) LoadAnswers(c *gin.Context, targetId string) ([]model.SubComment, error) {
	return cs.cd.LoadAnswers(c, targetId)
}

func toComment(c *gin.Context, req *req.CommentReq) *model.Comment {
	var cmt model.Comment
	cmt.CreatedAt = time.Now()
	cmt.Bid = tools.GenUUID(c)
	cmt.Content = req.Content
	cmt.CreatorID = req.CreatorID
	cmt.TargetID = req.TargetID
	return &cmt
}

func toAnswer(c *gin.Context, req *req.CommentReq) *model.SubComment {
	var cmt model.SubComment
	cmt.CreatedAt = time.Now()
	cmt.Bid = tools.GenUUID(c)
	cmt.Content = req.Content
	cmt.CreatorID = req.CreatorID
	cmt.TargetID = req.TargetID
	return &cmt
}
