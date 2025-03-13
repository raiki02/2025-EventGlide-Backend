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
	id *dao.InteractionDao
}

func NewCommentService(cd *dao.CommentDao, ud *dao.UserDao, id *dao.InteractionDao) *CommentService {
	return &CommentService{
		cd: cd,
		ud: ud,
		id: id,
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
	replys, err := cs.cd.LoadAnswers(c, cmt.Bid)
	if err != nil {
		return resp.CommentResp{}
	}
	res.Content = cmt.Content
	res.CommentedTime = cmt.CreatedAt.String()
	res.Bid = cmt.Bid
	res.CommentedPos = cmt.Position
	res.LikeNum = cmt.LikeNum
	res.ReplyNum = cmt.ReplyNum
	res.Creator.StudentID = user.StudentID
	res.Creator.Username = user.Name
	res.Creator.Avatar = user.Avatar
	for _, rep := range replys {
		cs.processReply(c, &res, &rep)
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

func (cs *CommentService) CreateComment(c *gin.Context, r req.CreateCommentReq) (resp.CommentResp, error) {
	cmt := cs.toComment(r)
	err := cs.cd.CreateComment(c, cmt)
	if err != nil {
		return resp.CommentResp{}, err
	}
	switch r.Subject {
	case "activity":
		err = cs.id.CommentActivity(c, r.StudentID, r.ParentID)
	case "post":
		err = cs.id.CommentPost(c, r.StudentID, r.ParentID)
	case "comment":
		err = cs.id.CommentComment(c, r.StudentID, r.ParentID)

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
	res := cs.toResps(c, cmts)
	return res, nil
}

func (cs *CommentService) processReply(c *gin.Context, r *resp.CommentResp, rep *model.Comment) {
	var tmp struct {
		Reply struct {
			Bid          string `json:"bid"`
			ReplyCreator struct {
				StudentID string `json:"studentid"`
				Username  string `json:"username"`
				Avatar    string `json:"avatar"`
			} `json:"reply_creator"`
			ReplyContent string `json:"reply_content"`
			ReplyTime    string `json:"reply_time"`
			ReplyPos     string `json:"reply_pos"`
		} `json:"reply"`
	}
	tmp.Reply.Bid = rep.Bid
	tmp.Reply.ReplyTime = rep.CreatedAt.String()
	tmp.Reply.ReplyContent = rep.Content
	tmp.Reply.ReplyPos = rep.Position
	user, err := cs.ud.GetUserInfo(c, rep.StudentID)
	if err != nil {
		return
	}
	tmp.Reply.ReplyCreator.StudentID = user.StudentID
	tmp.Reply.ReplyCreator.Username = user.Name
	tmp.Reply.ReplyCreator.Avatar = user.Avatar
	r.Reply = append(r.Reply, tmp.Reply)
}
