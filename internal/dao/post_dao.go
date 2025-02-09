package dao

import (
	"context"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type PostDaoHdl interface {
	GetAllPost(context.Context) ([]model.Post, error)
	CreatePost(context.Context, *model.Post) error
	FindPostByName(context.Context, string) ([]model.Post, error)
	DeletePost(context.Context, *model.Post) error
}

type PostDao struct {
	db *gorm.DB
}

func NewPostDao(db *gorm.DB) PostDaoHdl {
	return &PostDao{
		db: db,
	}
}

func (pd *PostDao) GetAllPost(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	err := pd.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pd *PostDao) CreatePost(ctx context.Context, post *model.Post) error {
	err := pd.db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (pd *PostDao) FindPostByName(ctx context.Context, name string) ([]model.Post, error) {
	var posts []model.Post
	err := pd.db.Where("name = ?", "%name%").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pd *PostDao) DeletePost(ctx context.Context, post *model.Post) error {
	err := pd.db.Delete(post).Error
	if err != nil {
		return err
	}
	return nil
}
