package dao

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type ActDaoHdl interface {
	CreateAct(*gin.Context, model.Activity) error
	CreateDraft(*gin.Context, model.ActivityDraft) error
	DeleteAct(*gin.Context, model.Activity) error
	LoadDraft(*gin.Context, string, string) (*model.ActivityDraft, error)
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

func (ad ActDao) CreateAct(c *gin.Context, a *model.Activity) error {
	if ad.CheckExist(c, a) {
		return errors.New("activity exist")
	} else {
		return ad.db.Create(a).Error
	}
}

func (ad ActDao) CreateDraft(c *gin.Context, d model.ActivityDraft) error {
	return ad.db.Create(&d).Error
}

func (ad ActDao) LoadDraft(c *gin.Context, s string, b string) (*model.ActivityDraft, error) {
	var d model.ActivityDraft
	err := ad.db.Where("creator_id = ? and bid = ?", s, b).Find(&d).Error
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// TODO: 是否换成按页展示，每页返回固定个数活动

func (ad ActDao) FindActByUser(c *gin.Context, s string, keyword string) ([]model.Activity, error) {
	var as []model.Activity
	if keyword == "" {
		err := ad.db.Where("creator_id = ? ", s).Find(&as).Error
		if err != nil {
			return nil, err
		}
		return as, nil
	} else {
		err := ad.db.Where("creator_id = ? and name like ?", s, fmt.Sprintf("%%%s%%", keyword)).Find(&as).Error
		if err != nil {
			return nil, err
		}
		return as, nil
	}
}

func (ad ActDao) FindActByName(c *gin.Context, n string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("name like ?", fmt.Sprintf("%%%s%%", n)).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad ActDao) FindActByDate(c *gin.Context, d string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("start_time like ?", fmt.Sprintf("%%%s%%", d)).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad ActDao) CheckExist(c *gin.Context, a *model.Activity) bool {
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

// todo 过滤器
func (ad ActDao) FindActBySearches(c *gin.Context, a *model.Activity) ([]model.Activity, error) {
	var as []model.Activity
	h := fmt.Sprintf("%%%s%%", a.Host)
	l := fmt.Sprintf("%%%s%%", a.Location)
	t := fmt.Sprintf("%%%s%%", a.Type)
	st := fmt.Sprintf("%%%s%%", a.StartTime)
	et := fmt.Sprintf("%%%s%%", a.EndTime)
	ir := fmt.Sprintf("%%%s%%", a.IfRegister)
	err := ad.db.Where("host like ? and location like ? and type like ? and start_time like ? and end_time like ? and if_register like ?", h, l, t, st, et, ir).Find(&as).Error
	//err := ad.db.Where(&a).Find(&as).Error
	return as, err
}
