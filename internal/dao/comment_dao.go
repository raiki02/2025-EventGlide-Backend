package dao

import (
	"context"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type CommentDaoHdl interface {
	Create(context.Context, model.Comment) error
}

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) *CommentDao {
	return &CommentDao{db}
}

func (dao *CommentDao) Create(ctx context.Context, comment model.Comment) error {
	return dao.db.Create(&comment).Error
}
