package dao

import (
	"context"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type UserDAOHdl interface {
	UpdateAvatar(context.Context) error
	UpdateUsername(context.Context) error
	Create(context.Context, string) error
	CheckUserExist(context.Context, int) bool
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAOHdl {
	return &UserDAO{db: db}
}

func (dao *UserDAO) UpdateAvatar(ctx context.Context) error {
	return nil
}

func (dao *UserDAO) UpdateUsername(ctx context.Context) error {
	return nil
}

// 新建用户时默认username是studentid，默认头像全一样/头像库随机
func (dao *UserDAO) Create(ctx context.Context, ssid string) error {

}

func (dao *UserDAO) CheckUserExist(ctx context.Context, sid int) bool {

}
