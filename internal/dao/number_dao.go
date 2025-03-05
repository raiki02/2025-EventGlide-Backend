package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type NumberDaoHdl interface {
	Insert(*model.Number) error
	Delete(*gin.Context, string, string) error
	Search(*gin.Context, string, string, string) ([]*model.Number, int, error)
	Update()
}

type NumberDao struct {
	db *gorm.DB
}

func NewNumberDao(db *gorm.DB) *NumberDao {
	return &NumberDao{
		db: db,
	}
}

func (nd *NumberDao) Insert(n *model.Number) error {
	return nd.db.Create(n).Error
}

func (nd *NumberDao) Delete(c *gin.Context, sid, obj string) error {
	return nd.db.Where("to_sid = ? and object = ? and action = like", sid, obj).Delete(&model.Number{}).Error
}

func (nd *NumberDao) Search(c *gin.Context, sid, obj, act string) ([]*model.Number, int, error) {
	var numbers []*model.Number
	var count int64
	err := nd.db.Where("to_sid = ? and object = ? and action = ? and is_read = 0", sid, obj, act).Find(&numbers).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	nd.db.Where("to_sid = ? and object = ? and action = ? and is_read = 0", sid, obj, act).Update("is_read", 1)

	return numbers, int(count), nil
}

func (nd *NumberDao) Update() {

}