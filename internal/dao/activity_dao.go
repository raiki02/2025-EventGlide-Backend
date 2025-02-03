package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type ActDaoHdl interface {
	CreateAct(*gin.Context, *model.Activity) error
	CreateDraft(*gin.Context, *model.ActivityDraft) error

	FindAllActs(*gin.Context) ([]model.Activity, error)
	FindActByHost(*gin.Context, string) ([]model.Activity, error)
	FindActByType(*gin.Context, string) ([]model.Activity, error)
	FindActByLocation(*gin.Context, string) ([]model.Activity, error)
	FindActByIfSignup(*gin.Context, string) ([]model.Activity, error)
	FindActByIsForeign(*gin.Context, string) ([]model.Activity, error)

	FindActByTime(*gin.Context, string, string) ([]model.Activity, error)
	FindActByName(*gin.Context, string) ([]model.Activity, error)

	CheckExist(*gin.Context, *model.Activity) bool
}

type ActDao struct {
	db *gorm.DB
}

func NewActDao(db *gorm.DB) ActDaoHdl {
	return &ActDao{
		db: db,
	}
}

func (ad ActDao) CreateAct(c *gin.Context, a *model.Activity) error {
	return nil
}
func (ad ActDao) CreateDraft(c *gin.Context, d *model.ActivityDraft) error {
	return nil
}

// TODO: 是否换成按页展示，每页返回固定个数活动
func (ad ActDao) FindAllActs(c *gin.Context) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Find(as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad ActDao) FindActByHost(c *gin.Context, h string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("host = ? ", h).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}
func (ad ActDao) FindActByType(c *gin.Context, t string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("type = ? ", t).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil

}
func (ad ActDao) FindActByLocation(c *gin.Context, l string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("location = ? ", l).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil

}
func (ad ActDao) FindActByIfSignup(c *gin.Context, sn string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("if_register = ? ", sn).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil

}
func (ad ActDao) FindActByIsForeign(c *gin.Context, f string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("visibility = ? ", f).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil

}

func (ad ActDao) FindActByTime(c *gin.Context, start string, end string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("start_time >= ? and end_time <= ?", start, end).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}
func (ad ActDao) FindActByName(c *gin.Context, n string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("name like ?", "%n%").Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil

}

func (ad ActDao) CheckExist(c *gin.Context, a *model.Activity) bool {

	return nil
}
