package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type CommentDaoHdl interface {
	CreateComment(*gin.Context, *model.Comment) error
	DeleteComment(*gin.Context, string) error
	AnswerComment(*gin.Context, *model.SubComment) error
}

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) *CommentDao {
	return &CommentDao{
		db: db,
	}
}

func (cd *CommentDao) CreateComment(c *gin.Context, cmt *model.Comment) error {
	return cd.db.Create(cmt).Error
}

func (cd *CommentDao) DeleteComment(c *gin.Context, sid string, targetID string) error {

	return cd.db.Where("creator_id = ? and target_id = ?", sid, targetID).Delete(&model.Comment{}).Error
}

func (cd *CommentDao) AnswerComment(c *gin.Context, cmt *model.SubComment) error {
	return cd.db.Create(cmt).Error
}
