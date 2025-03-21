package dao

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ActDaoHdl interface {
	CreateAct(*gin.Context, *model.Activity) error
	CreateDraft(*gin.Context, *model.ActivityDraft) error
	DeleteAct(*gin.Context, model.Activity) error
	LoadDraft(*gin.Context, string, string) (*model.ActivityDraft, error)
	FindActByName(*gin.Context, string) ([]model.Activity, error)
	FindActByDate(*gin.Context, string) ([]model.Activity, error)
	FindActByOwnerID(*gin.Context, string) ([]model.Activity, error)
	CheckExist(*gin.Context, *model.Activity) bool
	ListAllActs(*gin.Context) ([]model.Activity, error)
	FindActBySearches(*gin.Context, *req.ActSearchReq) ([]model.Activity, error)
}

type ActDao struct {
	db *gorm.DB
	l  *zap.Logger
}

func NewActDao(db *gorm.DB, l *zap.Logger) *ActDao {
	return &ActDao{
		db: db,
		l:  l.Named("activity/dao"),
	}
}

func (ad *ActDao) CreateAct(c *gin.Context, a *model.Activity) error {
	if ad.CheckExist(c, a) {
		ad.l.Warn("tried to create an exist activity", zap.Any("act-bid", a.Bid))
		return errors.New("activity exist")
	} else {
		ad.db.Where("student_id = ?", a.StudentID).Delete(model.ActivityDraft{})
		return ad.db.Create(a).Error
	}
}

func (ad *ActDao) CreateDraft(c *gin.Context, d *model.ActivityDraft) error {
	ad.db.Where("student_id = ?", d.StudentID).Delete(&model.ActivityDraft{})
	return ad.db.Create(d).Error
}

func (ad *ActDao) LoadDraft(c *gin.Context, s string) (model.ActivityDraft, error) {
	var d model.ActivityDraft
	err := ad.db.Where("student_id = ?", s).Find(&d).Error
	if err != nil {
		return model.ActivityDraft{}, err
	}
	return d, nil
}

// TODO: 换成按页展示，每页返回固定个数活动
func (ad *ActDao) FindActByUser(c *gin.Context, s string, keyword string) ([]model.Activity, error) {
	var as []model.Activity
	if keyword == "" {
		err := ad.db.Where("student_id = ? ", s).Find(&as).Error
		if err != nil {
			return nil, err
		}
		return as, nil
	} else {
		err := ad.db.Where("student_id = ? and title like ?", s, fmt.Sprintf("%%%s%%", keyword)).Find(&as).Error
		if err != nil {
			return nil, err
		}
		return as, nil
	}
}

func (ad *ActDao) FindActByName(c *gin.Context, n string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("title like ?", fmt.Sprintf("%%%s%%", n)).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad *ActDao) FindActByDate(c *gin.Context, d string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("start_time like ?", fmt.Sprintf("%%%s%%", d)).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad *ActDao) CheckExist(c *gin.Context, a *model.Activity) bool {
	ret := ad.db.Where(&model.Activity{
		Type:       a.Type,
		HolderType: a.HolderType,
		Position:   a.Position,
		IfRegister: a.IfRegister,
	}).Find(&model.Activity{}).RowsAffected
	if ret == 0 {
		return false
	} else {
		return true
	}
}

func (ad *ActDao) DeleteAct(c *gin.Context, a model.Activity) error {
	ret := ad.db.Where(&model.Activity{
		Type:       a.Type,
		HolderType: a.HolderType,
		Position:   a.Position,
		IfRegister: a.IfRegister,
	}).Find(&model.Activity{}).Delete(&model.Activity{}).RowsAffected
	if ret == 0 {
		return errors.New("activity not exist")
	} else {
		return nil
	}
}

func (ad *ActDao) FindActBySearches(c *gin.Context, a *req.ActSearchReq) ([]model.Activity, error) {
	var as []model.Activity
	q := ad.db // 确保 q 初始化
	if len(a.Type) > 0 {
		q = q.Where("type IN ?", a.Type)
	}
	if len(a.HolderType) > 0 {
		q = q.Where("holder_type IN ?", a.HolderType)
	}
	if len(a.Location) > 0 {
		q = q.Where("position IN ?", a.Location)
	}
	if a.IfRegister != "" {
		q = q.Where("if_register = ?", a.IfRegister)
	}
	if a.DetailTime.StartTime != "" && a.DetailTime.EndTime != "" {
		q = q.Where("start_time >= ? AND end_time <= ?", a.DetailTime.StartTime, a.DetailTime.EndTime)
	}

	err := q.Find(&as).Error
	return as, err
}

func (ad *ActDao) FindActByOwnerID(c *gin.Context, s string) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Where("student_id = ?", s).Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad *ActDao) ListAllActs(c *gin.Context) ([]model.Activity, error) {
	var as []model.Activity
	err := ad.db.Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (ad *ActDao) FindActByBid(c *gin.Context, bid string) (model.Activity, error) {
	var act model.Activity
	err := ad.db.Where("bid = ?", bid).Find(&act).Error
	if err != nil {
		return model.Activity{}, err
	}
	return act, nil
}
