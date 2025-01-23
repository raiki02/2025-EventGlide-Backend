package dao

import (
	"context"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type UserDAOHdl interface {
	UpdateAvatar(context.Context) error
	UpdateUsername(context.Context) error
	Insert(context.Context, string, string) error
	CheckUserExist(context.Context, int) bool
	FindUserById(context.Context, string) (model.User, error)
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
func (dao *UserDAO) Insert(ctx context.Context, ssid, pwd string) error {
	u := model.User{
		Name:      ssid,
		StudentId: ssid,
		Avatar:    model.DefaultAvatar,
	}
	return dao.db.Create(&u).Error
}

// 检查是否存在
func (dao *UserDAO) CheckUserExist(ctx context.Context, sid int) bool {
	var u model.User
	dao.db.Where("student_id = ?", sid).First(&u)
	return u.Id != 0
}

// 检查还要返回
func (dao *UserDAO) FindUserById(ctx context.Context, sid string) (model.User, error) {
	var u model.User
	err := dao.db.Where("student_id = ?", sid).First(&u).Error
	return u, err
}
