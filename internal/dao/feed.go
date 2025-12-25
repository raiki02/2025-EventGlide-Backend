package dao

import (
	"context"
	"errors"
	"fmt"
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
	var existing model.Feed
	if fd.db.WithContext(ctx).Where("receiver = ? AND student_id = ? AND action = ? AND object = ? AND target_bid = ? AND status = ?", feed.Receiver, feed.StudentId, feed.Action, feed.Object, feed.TargetBid, feed.Status).First(&existing); existing.Id != 0 {
		// 已经存在
		fmt.Println("feed重复操作，忽略创建")
		return nil
	}
	return fd.db.Create(feed).Error
}

func (fd *FeedDao) GetTotalCnt(ctx *gin.Context, id string) ([]int, error) {
	var lc, ca int64
	err1 := fd.db.Model(&model.Feed{}).Where("receiver = ? and action in ? and status = ? and student_id != ?", id, []string{"like", "collect"}, "未读", id).Count(&lc).Error
	err2 := fd.db.Model(&model.Feed{}).Where("receiver = ? and action in ? and status = ? and student_id != ?", id, []string{"comment", "at"}, "未读", id).Count(&ca).Error
	if err1 != nil || err2 != nil {
		fd.l.Error("Get Total Cnt Failed", zap.Error(err1), zap.Error(err2))
		return nil, errors.Join(err1, err2)
	}
	return []int{int(lc), int(ca), int(lc + ca)}, nil
}

func (fd *FeedDao) GetLikeFeed(ctx *gin.Context, id string) ([]*model.Feed, error) {
	var feeds []*model.Feed
	err := fd.db.Where("receiver = ? and action = ? and student_id != ?", id, "like", id).Find(&feeds).Error
	if err != nil {
		fd.l.Error("Get Like Feed Failed", zap.Error(err))
		return nil, err
	}
	return feeds, nil
}

func (fd *FeedDao) GetCollectFeed(ctx *gin.Context, id string) ([]*model.Feed, error) {
	var feeds []*model.Feed
	err := fd.db.Where("receiver = ? and action = ? and student_id != ?", id, "collect", id).Find(&feeds).Error
	if err != nil {
		fd.l.Error("Get Collect Feed Failed", zap.Error(err))
		return nil, err
	}
	return feeds, nil
}

func (fd *FeedDao) GetCommentFeed(ctx *gin.Context, id string) ([]*model.Feed, error) {
	var feeds []*model.Feed
	err := fd.db.Where("receiver = ? and action = ? and student_id != ?", id, "comment", id).Find(&feeds).Error
	if err != nil {
		fd.l.Error("Get Comment Feed Failed", zap.Error(err))
		return nil, err
	}
	return feeds, nil
}

func (fd *FeedDao) GetAtFeed(ctx *gin.Context, id string) ([]*model.Feed, error) {
	var feeds []*model.Feed
	err := fd.db.Where("receiver = ? and action = ? and student_id != ?", id, "at", id).Find(&feeds).Error
	if err != nil {
		fd.l.Error("Get At Feed Failed", zap.Error(err))
		return nil, err
	}
	return feeds, nil
}

func (fd *FeedDao) GetAuditorFeed(ctx *gin.Context, id string) ([]*model.Approvement, error) {
	var a []*model.Approvement
	if err := fd.db.WithContext(ctx).Where("stance = ? and student_id = ?", "pending", id).Find(&a).Error; err != nil {
		fd.l.Error("Get Auditor Feed Failed", zap.Error(err))
		return nil, err
	}
	return a, nil
}

func (fd *FeedDao) ReadFeedDetail(ctx *gin.Context, sid, id string) error {
	return fd.db.WithContext(ctx).Model(&model.Feed{}).Where("receiver = ? AND id = ? ", sid, id).Update("status", "已读").Error
}

func (fd *FeedDao) ReadAllFeed(ctx *gin.Context, sid string) error {
	return fd.db.WithContext(ctx).Model(&model.Feed{}).Where("receiver = ? ", sid).Update("status", "已读").Error
}

func (fd *FeedDao) GetPictureFromObj(ctx *gin.Context, targetId, object string) (string, error) {
	type Result struct {
		ShowImg string `gorm:"column:show_img"`
	}
	var res Result
	err := fd.db.WithContext(ctx).Table(object).Where("bid = ?", targetId).Select("show_img").Find(&res).Error
	if err != nil {
		fd.l.Error("Get First Pic Failed", zap.Error(err))
		return "", err
	}

	return res.ShowImg, nil
}
