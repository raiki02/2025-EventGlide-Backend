package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/api/resp"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/internal/mq"
	"github.com/raiki02/EG/tools"
	"go.uber.org/zap"
	"strings"
	"time"
)

type CommentServiceHdl interface {
}

type CommentService struct {
	cd *dao.CommentDao
	ud *dao.UserDao
	id *dao.InteractionDao
	mq *mq.MQ
	l  *zap.Logger
}

func NewCommentService(cd *dao.CommentDao, ud *dao.UserDao, id *dao.InteractionDao, l *zap.Logger, mq *mq.MQ) *CommentService {
	return &CommentService{
		cd: cd,
		ud: ud,
		id: id,
		mq: mq,
		l:  l.Named("comment/service"),
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

func (cs *CommentService) CreateComment(c *gin.Context, r req.CreateCommentReq) (resp.CommentResp, error) {
	cmt := cs.toComment(r)
	err := cs.cd.CreateComment(c, cmt)
	cs.l.Info("CreateComment",
		zap.String("bid", cmt.Bid),
		zap.String("studentid", cmt.StudentID),
		zap.String("parentid", cmt.ParentID),
	)

	if err != nil {
		cs.l.Error("comment create failed", zap.Error(err))
		return resp.CommentResp{}, err
	}

	f := model.Feed{
		StudentId: r.StudentID,
		TargetBid: r.ParentID,
		Object:    r.Subject,
		Action:    "comment",
	}
	err = cs.mq.Publish(c.Request.Context(), "feed", f)
	if err != nil {
		cs.l.Error("Publish Comment Feed Failed", zap.Error(err))
	}
	switch r.Subject {
	case "activity":
		err = cs.id.CommentActivity(c, r.StudentID, r.ParentID)
	case "post":
		err = cs.id.CommentPost(c, r.StudentID, r.ParentID)
	case "comment":
		err = cs.id.CommentComment(c, r.StudentID, r.ParentID)
	}
	if err != nil {
		cs.l.Error("comment create failed", zap.Error(err))
		return resp.CommentResp{}, err
	}

	return cs.toResp(c, cmt), nil
}

func (cs *CommentService) DeleteComment(c *gin.Context, r req.DeleteCommentReq) error {
	err := cs.cd.DeleteComment(c, r.StudentID, r.TargetID)
	if err != nil {
		cs.l.Error("comment delete failed", zap.Error(err))
		return err
	}
	return nil
}

func (cs *CommentService) AnswerComment(c *gin.Context, r req.CreateCommentReq) (resp.ReplyResp, error) {
	cmt := cs.toComment(r)
	err := cs.cd.AnswerComment(c, cmt)
	if err != nil {
		cs.l.Error("comment answer failed", zap.Error(err))
		return resp.ReplyResp{}, err
	}
	cs.l.Info("AnswerComment",
		zap.String("bid", cmt.Bid),
		zap.String("studentid", cmt.StudentID),
	)
	f := model.Feed{
		StudentId: r.StudentID,
		TargetBid: r.ParentID,
		Object:    "comment",
		Action:    "at",
	}
	err = cs.mq.Publish(c.Request.Context(), "feed", f)
	if err != nil {
		cs.l.Error("Publish Comment Feed Failed", zap.Error(err))
	}

	return cs.toReply(c, cmt), nil
}

func (cs *CommentService) LoadComments(c *gin.Context, parentid string) ([]resp.CommentResp, error) {
	cmts, err := cs.cd.LoadComments(c, parentid)
	if err != nil {
		cs.l.Error("load comments failed", zap.Error(err))
		return nil, err
	}
	res := cs.toResps(c, cmts)
	return res, nil
}

func (cs *CommentService) toResp(c *gin.Context, cmt *model.Comment) resp.CommentResp {
	var res resp.CommentResp
	user, err := cs.ud.GetUserInfo(c, cmt.StudentID)
	if err != nil {
		cs.l.Error("get user info when comment to resp", zap.Error(err))
		return resp.CommentResp{}
	}
	replys, err := cs.cd.LoadAnswers(c, cmt.Bid)
	if err != nil {
		cs.l.Error("load answers when loading replies", zap.Error(err))
		return resp.CommentResp{}
	}
	if strings.Contains(user.LikeComment, cmt.Bid) {
		res.IsLike = "true"
	} else {
		res.IsLike = "false"
	}
	res.Content = cmt.Content
	res.CommentedTime = tools.ParseTime(cmt.CreatedAt)
	res.Bid = cmt.Bid
	res.CommentedPos = cmt.Position
	res.LikeNum = cmt.LikeNum
	res.ReplyNum = cmt.ReplyNum
	res.Creator.StudentID = user.StudentID
	res.Creator.Username = user.Name
	res.Creator.Avatar = user.Avatar
	for _, reply := range replys {
		res.Reply = append(res.Reply, cs.toReply(c, &reply))
	}
	return res
}

func (cs *CommentService) toResps(c *gin.Context, cmts []model.Comment) []resp.CommentResp {
	var res []resp.CommentResp
	for _, cmt := range cmts {
		res = append(res, cs.toResp(c, &cmt))
	}
	return res
}

func (cs *CommentService) toReply(c *gin.Context, cmt *model.Comment) resp.ReplyResp {
	var res resp.ReplyResp
	user, err := cs.ud.GetUserInfo(c, cmt.StudentID)
	if err != nil {
		cs.l.Error("get user info when comment to reply", zap.Error(err))
		return resp.ReplyResp{}
	}
	pid := cmt.ParentID
	pc := cs.cd.FindCmtByID(c, pid)
	if pc == nil {
		cs.l.Error("find comment by id", zap.String("pid", pid))
		return resp.ReplyResp{}
	}
	pu, err := cs.ud.GetUserInfo(c, pc.StudentID)
	if err != nil {
		cs.l.Error("get user info when comment to reply", zap.Error(err))
		return resp.ReplyResp{}
	}
	res.ReplyContent = cmt.Content
	res.ReplyTime = tools.ParseTime(cmt.CreatedAt)
	res.Bid = cmt.Bid
	res.ReplyPos = cmt.Position
	res.LikeNum = cmt.LikeNum
	res.ReplyNum = cmt.ReplyNum
	res.ReplyCreator.StudentID = user.StudentID
	res.ReplyCreator.Username = user.Name
	res.ReplyCreator.Avatar = user.Avatar
	res.ParentUserName = pu.Name
	return res
}
