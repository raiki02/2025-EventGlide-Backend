package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type UserDaoHdl interface {
	UpdateAvatar(*gin.Context, string, string) error
	UpdateUsername(*gin.Context, string, string) error
	Create(*gin.Context, *model.User) error
	CheckUserExist(*gin.Context, string) bool
	GetUserInfo(*gin.Context, string) (model.User, error)
	FindUserByID(*gin.Context, string) model.User
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (ud *UserDao) UpdateAvatar(ctx *gin.Context, student_id string, imgurl string) error {
	return ud.db.Model(&model.User{}).Where("student_id = ?", student_id).Update("avatar", imgurl).Error
}

func (ud *UserDao) UpdateUsername(ctx *gin.Context, student_id string, name string) error {
	return ud.db.Model(&model.User{}).Where("student_id = ?", student_id).Update("name", name).Error
}

func (ud *UserDao) Create(ctx *gin.Context, user *model.User) error {
	return ud.db.Create(user).Error
}

func (ud *UserDao) CheckUserExist(ctx *gin.Context, student_id string) bool {
	res := ud.db.Where("student_id = ?", student_id).Find(&model.User{}).RowsAffected
	return res != 0
}

func (ud *UserDao) GetUserInfo(ctx *gin.Context, student_id string) (model.User, error) {
	var user model.User
	err := ud.db.Where("student_id = ?", student_id).First(&user).Error
	return user, err
}

func (ud *UserDao) FindUserByID(ctx *gin.Context, student_id string) model.User {
	var user model.User
	err := ud.db.Where("student_id = ?", student_id).First(&user).Error
	if err != nil {
		return model.User{}
	}
	return user
}
