package dao

import (
	"context"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type CommentDAOHdl interface {
	Create(context.Context, model.Comment) error
}

type CommentDAO struct {
	db *gorm.DB
}

func NewCommentDAO(db *gorm.DB) *CommentDAO {
	return &CommentDAO{db}
}

func (dao *CommentDAO) Create(ctx context.Context, comment model.Comment) error {
	return dao.db.Create(&comment).Error
}
