package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type NumberDaoHdl interface {
	AddLikesNum(*gin.Context, *model.Number) error
	CutLikesNum(*gin.Context, string, string) error
	AddCommentsNum(*gin.Context, *model.Number) error
}

type NumberDao struct {
	db *gorm.DB
}

func NewNumberDao(db *gorm.DB) *NumberDao {
	return &NumberDao{
		db: db,
	}
}

func (nd *NumberDao) AddLikesNum(c *gin.Context, n *model.Number) error {
	var num model.Number
	if nd.db.Where("sid = ? AND bid = ? AND topic = ?", n.Sid, n.Bid, "like").First(&num).RowsAffected > 0 {
		return errors.New("already liked")
	}
	return nd.db.Create(n).Error
}

func (nd *NumberDao) CutLikesNum(c *gin.Context, sid string, bid string) error {
	return nd.db.Where("sid = ? AND bid = ? AND topic = ?", sid, bid, "like").Delete(&model.Number{}).Error
}

func (nd *NumberDao) AddCommentsNum(c *gin.Context, n *model.Number) error {
	return nd.db.Create(n).Error
}

func (nd *NumberDao) GetLikesNum(c *gin.Context, sid string, bid string) int {
	var num int64
	nd.db.Model(&model.Number{}).Where("sid = ? AND bid = ? AND topic = ?", sid, bid, "like").Count(&num)
	return int(num)
}

func (nd *NumberDao) GetCommentsNum(c *gin.Context, sid string, bid string) int {
	var num int64
	nd.db.Model(&model.Number{}).Where("sid = ? AND bid = ? AND topic = ?", sid, bid, "comment").Count(&num)
	return int(num)
}

//todo check bid exist
