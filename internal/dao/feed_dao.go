package dao

import (
	"gorm.io/gorm"
)

type NumberDaoHdl interface {
}

type NumberDao struct {
	db *gorm.DB
}

func NewNumberDao(db *gorm.DB) *NumberDao {
	return &NumberDao{
		db: db,
	}
}
