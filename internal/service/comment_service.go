package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/api/resp"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
	"time"
)

type CommentServiceHdl interface {
}

type CommentService struct {
	cd *dao.CommentDao
	ud *dao.UserDao
}

func NewCommentService(cd *dao.CommentDao, ud *dao.UserDao) *CommentService {
	return &CommentService{
		cd: cd,
		ud: ud,
	}
}

func (cs *CommentService) toComment(r req.CreateCommentReq) *model.Comment {
	return &model.Comment{
		StudentID: r.StudentID,
		Content:   r.Content,
		ParentID:  r.ParentID,
		CreatedAt: time.Now(),
		Bid:       tools.GenUUID(),
		Position:  "华中师范大学",
	}
}

func (cs *CommentService) toResp(c *gin.Context, cmt *model.Comment) resp.CommentResp {
	var res resp.CommentResp
	user, err := cs.ud.GetUserInfo(c, cmt.StudentID)
	if err != nil {
		return resp.CommentResp{}
	}
	res.Content = cmt.Content
	res.CommentedTime = cmt.CreatedAt.String()
	res.CommentedPos = cmt.Position
	res.LikeNum = cmt.LikeNum
	res.ReplyNum = cmt.ReplyNum
	res.Creator.StudentID = user.StudentID
	res.Creator.Username = user.Name
	res.Creator.Avatar = user.Avatar
	return res
}

func (cs *CommentService) toResps(c *gin.Context, cmts []model.Comment) []resp.CommentResp {
	var res []resp.CommentResp
	for _, cmt := range cmts {
		res = append(res, cs.toResp(c, &cmt))
	}
	return res
}

func (cs *CommentService) CreateComment(c *gin.Context, r req.CreateCommentReq) (resp.CommentResp, error) {
	cmt := cs.toComment(r)
	err := cs.cd.CreateComment(c, cmt)
	if err != nil {
		return resp.CommentResp{}, err
	}
	return cs.toResp(c, cmt), nil
}

func (cs *CommentService) DeleteComment(c *gin.Context, r req.DeleteCommentReq) error {
	err := cs.cd.DeleteComment(c, r.StudentID, r.TargetID)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CommentService) AnswerComment(c *gin.Context, r req.CreateCommentReq) (resp.CommentResp, error) {
	cmt := cs.toComment(r)
	err := cs.cd.AnswerComment(c, cmt)
	if err != nil {
		return resp.CommentResp{}, err
	}
	return cs.toResp(c, cmt), nil
}

func (cs *CommentService) LoadComments(c *gin.Context, parentid string) ([]resp.CommentResp, error) {
	cmts, err := cs.cd.LoadComments(c, parentid)
	if err != nil {
		return nil, err
	}
	return cs.toResps(c, cmts), nil
}

func (cs *CommentService) LoadAnswers(c *gin.Context, parentid string) ([]resp.CommentResp, error) {
	cmts, err := cs.cd.LoadAnswers(c, parentid)
	if err != nil {
		return nil, err
	}
	return cs.toResps(c, cmts), nil
}
