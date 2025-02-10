package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type UserDAOHdl interface {
	UpdateAvatar(*gin.Context, string, string) error
	UpdateUsername(*gin.Context, string, string) error
	Create(*gin.Context, *model.User) error
	CheckUserExist(*gin.Context, string) bool
	GetUserInfo(*gin.Context, string) (model.User, error)
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAOHdl {
	return &UserDAO{db: db}
}

func (ud *UserDAO) UpdateAvatar(ctx *gin.Context, sid string, imgurl string) error {
	return ud.db.Model(&model.User{}).Where("sid = ?", sid).Update("avatar", imgurl).Error
}

func (ud *UserDAO) UpdateUsername(ctx *gin.Context, sid string, name string) error {
	return ud.db.Model(&model.User{}).Where("sid = ?", sid).Update("name", name).Error
}

// 新建用户时默认username是studentid，默认头像全一样/头像库随机
func (ud *UserDAO) Create(ctx *gin.Context, user *model.User) error {
	return ud.db.Create(user).Error
}

func (ud *UserDAO) CheckUserExist(ctx *gin.Context, sid string) bool {
	res := ud.db.Where("sid = ?", sid).Find(&model.User{}).RowsAffected
	return res != 0
}

func (ud *UserDAO) GetUserInfo(ctx *gin.Context, sid string) (model.User, error) {
	var user model.User
	err := ud.db.Where("sid = ?", sid).First(&user).Error
	return user, err
}
