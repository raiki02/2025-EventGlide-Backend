package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type CommentDaoHdl interface {
}

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) *CommentDao {
	return &CommentDao{
		db: db,
	}
}

func (cd *CommentDao) CreateComment(c *gin.Context, cmt *model.Comment) error {
	return cd.db.Create(cmt).Error
}

func (cd *CommentDao) DeleteComment(c *gin.Context, sid, bid string) error {
	return cd.db.Where("student_id = ? and bid = ?", sid, bid).Delete(&model.Comment{}).Error
}

func (cd *CommentDao) AnswerComment(c *gin.Context, cmt *model.Comment) error {
	cmt.Type = 1
	return cd.db.Create(cmt).Error
}

func (cd *CommentDao) LoadComments(c *gin.Context, parentid string) ([]model.Comment, error) {
	var cmts []model.Comment
	err := cd.db.Where("parent_id = ? and type = 0", parentid).Find(&cmts).Error
	return cmts, err
}

func (cd *CommentDao) LoadAnswers(c *gin.Context, pid string) ([]model.Comment, error) {
	var cmts []model.Comment
	err := cd.db.Where("parent_id = ? and type = 1", pid).Find(&cmts).Error
	return cmts, err
}

func (cd *CommentDao) Like(c *gin.Context, bid string, t int) error {
	var cmt model.Comment
	switch t {
	case 1:
		return cd.db.Where("bid = ?", bid).First(&cmt).Update("like_num", cmt.LikeNum+1).Error
	case 0:
		return cd.db.Where("bid = ?", bid).First(&cmt).Update("like_num", cmt.LikeNum-1).Error
	default:
		return errors.New("type error")
	}
}
