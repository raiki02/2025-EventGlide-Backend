package dao

import (
	"context"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type ActivityDaoHdl interface {
	NewActivity(context.Context, model.Activity) error
	NewDraft(context.Context, model.ActivityDraft) error
	ListAllActivity(context.Context) ([]model.Activity, error)
	ListActivityByType(context.Context, string) ([]model.Activity, error)
	ListActivityByTime(context.Context, string) ([]model.Activity, error)
	ListActivityByHost(context.Context, string) ([]model.Activity, error)
	ListActivityByLocation(context.Context, string) ([]model.Activity, error)
	ListActivityByName(context.Context, string) ([]model.Activity, error)
}

type ActivityDao struct {
	db *gorm.DB
}

func NewActivityDao(db *gorm.DB) ActivityDaoHdl {
	return &ActivityDao{db: db}
}

func (a *ActivityDao) NewActivity(ctx context.Context, activity model.Activity) error {
	return a.db.Create(&activity).Error
}

func (a *ActivityDao) NewDraft(ctx context.Context, activityDraft model.ActivityDraft) error {
	return a.db.Create(&activityDraft).Error
}

func (a *ActivityDao) ListAllActivity(ctx context.Context) ([]model.Activity, error) {
	var activities []model.Activity
	err := a.db.Find(&activities).Error
	return activities, err
}

func (a *ActivityDao) ListActivityByType(ctx context.Context, t string) ([]model.Activity, error) {
	var activities []model.Activity
	err := a.db.Where("type = ?", t).Find(&activities).Error
	return activities, err
}

func (a *ActivityDao) ListActivityByTime(ctx context.Context, t string) ([]model.Activity, error) {
	var activities []model.Activity
	err := a.db.Where("time = ?", t).Find(&activities).Error
	return activities, err
}

func (a *ActivityDao) ListActivityByHost(ctx context.Context, h string) ([]model.Activity, error) {
	var activities []model.Activity
	err := a.db.Where("host = ?", h).Find(&activities).Error
	return activities, err
}

func (a *ActivityDao) ListActivityByLocation(ctx context.Context, l string) ([]model.Activity, error) {
	var activities []model.Activity
	err := a.db.Where("location = ?", l).Find(&activities).Error
	return activities, err
}

func (a *ActivityDao) ListActivityByName(ctx context.Context, n string) ([]model.Activity, error) {
	var activities []model.Activity
	err := a.db.Where("name = ?", n).Find(&activities).Error
	return activities, err
}
