package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FeedDaoHdl interface {
	CreateFeed(ctx context.Context, feed *model.Feed) error
	GetTotalCnt(ctx *gin.Context, id string) ([]int, error)
	GetLikeFeed(ctx *gin.Context, id string) ([]*model.Feed, error)
	GetCollectFeed(ctx *gin.Context, id string) ([]*model.Feed, error)
	GetCommentFeed(ctx *gin.Context, id string) ([]*model.Feed, error)
	GetAtFeed(ctx *gin.Context, id string) ([]*model.Feed, error)
	GetInvitationFeed(ctx *gin.Context, id string) ([]*model.Feed, error)
}

type FeedDao struct {
	db *gorm.DB
	l  *zap.Logger
}

func NewFeedDao(db *gorm.DB, l *zap.Logger) *FeedDao {
	return &FeedDao{
		db: db,
		l:  l.Named("feed/dao"),
	}
}

func (fd *FeedDao) CreateFeed(ctx context.Context, feed *model.Feed) error {
	return fd.db.Create(feed).Error
}

func (fd *FeedDao) GetTotalCnt(ctx *gin.Context, id string) ([]int, error) {
	var lc, ca int64
	err1 := fd.db.Model(&model.Feed{}).Where("student_id = ? and action in ? and status = ?", id, []string{"like", "collect"}, "未读").Count(&lc).Error
	err2 := fd.db.Model(&model.Feed{}).Where("student_id = ? and action in ? and status = ?", id, []string{"comment", "at"}, "未读").Count(&ca).Error
	if err1 != nil || err2 != nil {
		return nil, err1
	}
	return []int{int(lc), int(ca), int(lc + ca)}, nil
}

func (fd *FeedDao) GetLikeFeed(ctx *gin.Context, id string) ([]*model.Feed, error) {
	var feeds []*model.Feed
	err1 := fd.db.Where("student_id = ? and action = ?", id, "like").Find(&feeds).Error
	err2 := fd.db.Model(&model.Feed{}).Where("student_id = ? and action = ?", id, "like").Updates(map[string]interface{}{"status": "已读"}).Error
	if err1 != nil || err2 != nil {
		fd.l.Error("Get Like Feed Failed", zap.Error(err1), zap.Error(err2))
		return nil, err1
	}
	return feeds, nil
}

func (fd *FeedDao) GetCollectFeed(ctx *gin.Context, id string) ([]*model.Feed, error) {
	var feeds []*model.Feed
	err1 := fd.db.Where("student_id = ? and action = ?", id, "collect").Find(&feeds).Error
	err2 := fd.db.Model(&model.Feed{}).Where("student_id = ? and action = ?", id, "collect").Updates(map[string]interface{}{"status": "已读"}).Error
	if err1 != nil || err2 != nil {
		fd.l.Error("Get Collect Feed Failed", zap.Error(err1), zap.Error(err2))
		return nil, err1
	}
	return feeds, nil
}

func (fd *FeedDao) GetCommentFeed(ctx *gin.Context, id string) ([]*model.Feed, error) {
	var feeds []*model.Feed
	err1 := fd.db.Where("student_id = ? and action = ?", id, "comment").Find(&feeds).Error
	err2 := fd.db.Model(&model.Feed{}).Where("student_id = ? and action = ?", id, "comment").Updates(map[string]interface{}{"status": "已读"}).Error
	if err1 != nil || err2 != nil {
		fd.l.Error("Get Comment Feed Failed", zap.Error(err1), zap.Error(err2))
		return nil, err1
	}
	return feeds, nil
}

func (fd *FeedDao) GetAtFeed(ctx *gin.Context, id string) ([]*model.Feed, error) {
	var feeds []*model.Feed
	err1 := fd.db.Where("student_id = ? and action = ?", id, "at").Find(&feeds).Error
	err2 := fd.db.Model(&model.Feed{}).Where("student_id = ? and action = ?", id, "at").Updates(map[string]interface{}{"status": "已读"}).Error
	if err1 != nil || err2 != nil {
		fd.l.Error("Get At Feed Failed", zap.Error(err1), zap.Error(err2))
		return nil, err1
	}
	return feeds, nil
}

func (fd *FeedDao) GetAuditorFeed(ctx *gin.Context, id string) ([]*model.Approvement, error) {
	var a []*model.Approvement
	if err := fd.db.WithContext(ctx).Where("student_id = ? AND stance = ?", id, "pending").Find(&a).Error; err != nil {
		fd.l.Error("Get Auditor Feed Failed", zap.Error(err))
		return nil, err
	}
	return a, nil
}
