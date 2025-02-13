package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type ActDaoHdl interface {
	CreateAct(*gin.Context, model.Activity) error
	CreateDraft(*gin.Context, model.ActivityDraft) error
	DeleteAct(*gin.Context, model.Activity) error

	FindActByBid(*gin.Context, string) (model.Activity, error)
	FindActByName(*gin.Context, string) ([]model.Activity, error)
	FindActByDate(*gin.Context, string) ([]model.Activity, error)

	CheckExist(*gin.Context, model.Activity) bool

	FindActBySearches(*gin.Context, *model.Activity) ([]model.Activity, error)
}

type ActDao struct {
	db *gorm.DB
}

func NewActDao(db *gorm.DB) *ActDao {
	return &ActDao{
		db: db,
	}
}

func (ad ActDao) CreateAct(c *gin.Context, a model.Activity) error {
	if ad.CheckExist(c, a) {
		return errors.New("activity exist")
	} else {
		return ad.db.Create(&a).Error
	}
}

func (ad ActDao) CreateDraft(c *gin.Context, d model.ActivityDraft) error {
	return ad.db.Create(&d).Error
}

// TODO: 是否换成按页展示，每页返回固定个数活动

func (ad ActDao) FindActByUser(c *gin.Context, s string, keyword string) ([]model.Activity, error) {
	var as []model.Activity
	if keyword == "" {
		err := ad.db.Where("host = ? ", s).Find(&as).Error
		if err != nil {
			return nil, err
		}
		return as, nil
	} else {
		err := ad.db.Where("host = ? and name like ?", s, "%keyword%").Find(&as).Error
		if err != nil {
			return nil, err
		}
		return as, nil
	}
}

func (ad ActDao) FindActByName(c *gin.Context, n string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("name like ?", "%n%").Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad ActDao) FindActByDate(c *gin.Context, d string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("start_time like ?", "%d%").Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad ActDao) CheckExist(c *gin.Context, a model.Activity) bool {
	ret := ad.db.Where(&model.Activity{
		Type:       a.Type,
		Host:       a.Host,
		Location:   a.Location,
		IfRegister: a.IfRegister,
		Capacity:   a.Capacity,
	}).Find(&model.Activity{}).RowsAffected
	if ret == 0 {
		return false
	} else {
		return true
	}
}

func (ad ActDao) DeleteAct(c *gin.Context, a model.Activity) error {
	ret := ad.db.Where(&model.Activity{
		Type:       a.Type,
		Host:       a.Host,
		Location:   a.Location,
		IfRegister: a.IfRegister,
		Capacity:   a.Capacity,
	}).Find(&model.Activity{}).Delete(&model.Activity{}).RowsAffected
	if ret == 0 {
		return errors.New("activity not exist")
	} else {
		return nil
	}
}

func (ad ActDao) FindActBySearches(c *gin.Context, a *model.Activity) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where(&a).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad ActDao) FindActByBid(c *gin.Context, bid string) (model.Activity, error) {
	var a model.Activity
	err := ad.db.Where("bid = ?", bid).Find(&a).Error
	if err != nil {
		return model.Activity{}, err
	}
	return a, nil
}
