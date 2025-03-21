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
	mq *mq.MQ
	l  *zap.Logger
}

func NewInteractionService(id *dao.InteractionDao, mq *mq.MQ, l *zap.Logger) *InteractionService {
	return &InteractionService{
		id: id,
		mq: mq,
		l:  l.Named("interaction/service"),
	}
}

func (is *InteractionService) Like(c *gin.Context, r *req.InteractionReq) error {
	jreq := is.toFeed(r, "like")
	err := is.mq.Publish(c.Request.Context(), "feed", jreq)
	if err != nil {
		is.l.Error("Publish Like Feed Failed", zap.Error(err))
	}

	switch r.Subject {
	case "activity":
		return is.id.LikeActivity(c, r.StudentID, r.TargetID)
	case "post":
		return is.id.LikePost(c, r.StudentID, r.TargetID)
	case "comment":
		return is.id.LikeComment(c, r.StudentID, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) Dislike(c *gin.Context, r *req.InteractionReq) error {
	switch r.Subject {
	case "activity":
		return is.id.DislikeActivity(c, r.StudentID, r.TargetID)
	case "post":
		return is.id.DislikePost(c, r.StudentID, r.TargetID)
	case "comment":
		return is.id.DislikeComment(c, r.StudentID, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) Comment(c *gin.Context, r *req.InteractionReq) error {
	jreq := is.toFeed(r, "comment")
	err := is.mq.Publish(c.Request.Context(), "feed", jreq)
	if err != nil {
		is.l.Error("Publish Comment Feed Failed", zap.Error(err))
	}
	switch r.Subject {
	case "activity":
		return is.id.CommentActivity(c, r.StudentID, r.TargetID)
	case "post":
		return is.id.CommentPost(c, r.StudentID, r.TargetID)
	case "comment":
		return is.id.CommentComment(c, r.StudentID, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) Collect(c *gin.Context, r *req.InteractionReq) error {
	jreq := is.toFeed(r, "collect")
	err := is.mq.Publish(c.Request.Context(), "feed", jreq)
	if err != nil {
		is.l.Error("Publish Collect Feed Failed", zap.Error(err))
	}
	switch r.Subject {
	case "activity":
		return is.id.CollectActivity(c, r.StudentID, r.TargetID)
	case "post":
		return is.id.CollectPost(c, r.StudentID, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) Discollect(c *gin.Context, r *req.InteractionReq) error {
	switch r.Subject {
	case "activity":
		return is.id.DiscollectActivity(c, r.StudentID, r.TargetID)
	case "post":
		return is.id.DiscollectPost(c, r.StudentID, r.TargetID)
	default:
		return errors.New("subject error")
	}
}

func (is *InteractionService) toFeed(r *req.InteractionReq, action string) model.Feed {
	f := model.Feed{
		TargetBid: r.TargetID,
		Object:    r.Subject,
		StudentId: r.StudentID,
		Action:    action,
	}
	return f
}
