package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/dao"
)

type InteractionService struct {
	id *dao.InteractionDao
}

func NewInteractionService(id *dao.InteractionDao) *InteractionService {
	return &InteractionService{
		id: id,
	}
}

func (is *InteractionService) Like(c *gin.Context, r *req.InteractionReq) error {
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
