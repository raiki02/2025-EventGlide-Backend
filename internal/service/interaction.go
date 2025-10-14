package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/internal/mq"
	"go.uber.org/zap"
)

type InteractionService struct {
	id *dao.InteractionDao
	mq mq.MQHdl
	l  *zap.Logger
}

func NewInteractionService(id *dao.InteractionDao, mq mq.MQHdl, l *zap.Logger) *InteractionService {
	return &InteractionService{
		id: id,
		mq: mq,
		l:  l.Named("interaction/service"),
	}
}

func (is *InteractionService) Like(c *gin.Context, r *req.InteractionReq, sid string) error {
	recv, err := is.getReceiver(c, r)
	if err != nil {
		return err
	}
	jreq := is.toFeed(r, "like", sid, recv)

	err = is.mq.Publish(c.Request.Context(), "feed_stream", jreq)
	if err != nil {
		is.l.Error("Publish Like Feed Failed", zap.Error(err), zap.Any("feed", jreq))
	} else {
		is.l.Info("Publish Like Feed Success", zap.Any("feed", jreq))
	}

	switch r.Subject {
	case "activity":
		return is.id.LikeActivity(c, sid, r.TargetID)
	case "post":
		return is.id.LikePost(c, sid, r.TargetID)
	case "comment":
		return is.id.LikeComment(c, sid, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) Dislike(c *gin.Context, r *req.InteractionReq, sid string) error {
	switch r.Subject {
	case "activity":
		return is.id.DislikeActivity(c, sid, r.TargetID)
	case "post":
		return is.id.DislikePost(c, sid, r.TargetID)
	case "comment":
		return is.id.DislikeComment(c, sid, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) Comment(c *gin.Context, r *req.InteractionReq, sid string) error {
	recv, err := is.getReceiver(c, r)
	if err != nil {
		return err
	}

	jreq := is.toFeed(r, "comment", sid, recv)
	err = is.mq.Publish(c.Request.Context(), "feed_stream", jreq)
	if err != nil {
		is.l.Error("Publish Comment Feed Failed", zap.Error(err), zap.Any("feed", jreq))
	} else {
		is.l.Info("Publish Comment Feed Success", zap.Any("feed", jreq))
	}

	switch r.Subject {
	case "activity":
		return is.id.CommentActivity(c, sid, r.TargetID)
	case "post":
		return is.id.CommentPost(c, sid, r.TargetID)
	case "comment":
		return is.id.CommentComment(c, sid, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) Collect(c *gin.Context, r *req.InteractionReq, sid string) error {
	recv, err := is.getReceiver(c, r)
	if err != nil {
		return err
	}
	jreq := is.toFeed(r, "collect", sid, recv)
	err = is.mq.Publish(c.Request.Context(), "feed_stream", jreq)
	if err != nil {
		is.l.Error("Publish Collect Feed Failed", zap.Error(err), zap.Any("feed", jreq))
	} else {
		is.l.Info("Publish Collect Feed Success", zap.Any("feed", jreq))
	}

	switch r.Subject {
	case "activity":
		return is.id.CollectActivity(c, sid, r.TargetID)
	case "post":
		return is.id.CollectPost(c, sid, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) Discollect(c *gin.Context, r *req.InteractionReq, sid string) error {
	switch r.Subject {
	case "activity":
		return is.id.DiscollectActivity(c, sid, r.TargetID)
	case "post":
		return is.id.DiscollectPost(c, sid, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) toFeed(r *req.InteractionReq, action string, sid string, recv string) model.Feed {
	f := model.Feed{
		TargetBid: r.TargetID,
		Object:    r.Subject,
		StudentId: sid,
		Action:    action,
		Receiver:  recv,
	}
	return f
}

func (is *InteractionService) Approve(c *gin.Context, studendId string, r *req.InteractionReq) error {
	return is.id.ApproveActivity(c, studendId, r.TargetID)
}

func (is *InteractionService) Reject(c *gin.Context, studendId string, r *req.InteractionReq) error {
	return is.id.RejectActivity(c, studendId, r.TargetID)
}

func (is *InteractionService) getReceiver(c *gin.Context, r *req.InteractionReq) (string, error) {
	switch r.Subject {
	case "activity":
		act, err := is.id.GetActivityByBid(c, r.TargetID)
		if err != nil {
			is.l.Error("Get Activity Failed when get receiver", zap.Error(err), zap.String("bid", r.TargetID))
			return "", err
		}
		if act == nil {
			is.l.Error("Activity Not Found when get receiver", zap.String("bid", r.TargetID))
			return "", errors.New("activity not found")
		}
		return act.StudentID, nil
	case "post":
		post, err := is.id.GetPostByBid(c, r.TargetID)
		if err != nil {
			is.l.Error("Get Post Failed when get receiver", zap.Error(err), zap.String("bid", r.TargetID))
			return "", err
		}
		if post == nil {
			is.l.Error("Post Not Found when get receiver", zap.String("bid", r.TargetID))
			return "", errors.New("post not found")
		}
		return post.StudentID, nil
	case "comment":
		comment, err := is.id.GetCommentByBid(c, r.TargetID)
		if err != nil {
			is.l.Error("Get Comment Failed when get receiver", zap.Error(err), zap.String("bid", r.TargetID))
			return "", err
		}
		if comment == nil {
			is.l.Error("Comment Not Found when get receiver", zap.String("bid", r.TargetID))
			return "", errors.New("comment not found")
		}
		return comment.StudentID, nil
	default:
		return "", errors.New("subject error")
	}
}
