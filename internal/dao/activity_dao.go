package dao

import (
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type ActDaoHdl interface {
	CreateAct(*model.Activity) error
	CreateDraft(*model.ActivityDraft) error

	FindActByHost(string) error
	FindActByType(string) error
	FindActByLocation(string) error
	FindActByIfSignup(string) error
	FindActByIsForeign(string) error

	FindActByTime(string, string) error
	FindActByName(string) error

	CheckExist(*model.Activity) error
}

type ActDao struct {
	db *gorm.DB
}

func NewActDao(db *gorm.DB) ActDaoHdl {
	return &ActDao{
		db: db,
	}
}

func (ad ActDao) CreateAct(*model.Activity) error {
	return nil
}
func (ad ActDao) CreateDraft(*model.ActivityDraft) error {
	return nil

}

func (ad ActDao) FindActByHost(string) error {
	return nil

}
func (ad ActDao) FindActByType(string) error {
	return nil

}
func (ad ActDao) FindActByLocation(string) error {
	return nil

}
func (ad ActDao) FindActByIfSignup(string) error {
	return nil

}
func (ad ActDao) FindActByIsForeign(string) error {
	return nil

}

func (ad ActDao) FindActByTime(string, string) error {
	return nil

}
func (ad ActDao) FindActByName(string) error {
	return nil

}

func (ad ActDao) CheckExist(*model.Activity) error {
	return nil

}
