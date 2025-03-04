package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type CommentDaoHdl interface {
	CreateComment(*gin.Context, *model.Comment) error
	DeleteComment(*gin.Context, string) error
	AnswerComment(*gin.Context, *model.SubComment) error
	LoadComments(*gin.Context, string) ([]model.Comment, error)
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

func (cd *CommentDao) DeleteComment(c *gin.Context, sid string, targetID string) error {

	return cd.db.Where("creator_id = ? and target_id = ?", sid, targetID).Delete(&model.Comment{}).Error
}

func (cd *CommentDao) AnswerComment(c *gin.Context, cmt *model.SubComment) error {
	return cd.db.Create(cmt).Error
}

func (cd *CommentDao) UpdateNumbersForComments(c *gin.Context, sid, bid string, like, comment int) error {
	return cd.db.Model(&model.Comment{}).Where("creator_id = ? AND bid = ?", sid, bid).Update("like", like).Update("comment", comment).Error
}

func (cd *CommentDao) UpdateNumbersForAnswers(c *gin.Context, sid, bid string, like int) error {
	return cd.db.Model(&model.SubComment{}).Where("creator_id = ? AND bid = ?", sid, bid).Update("like", like).Error
}

func (cd *CommentDao) LoadComments(c *gin.Context, targetID string) ([]model.Comment, error) {
	var cmts []model.Comment
	err := cd.db.Where("target_id = ?", targetID).Find(&cmts).Error
	return cmts, err
}

func (cd *CommentDao) LoadAnswers(c *gin.Context, targetID string) ([]model.SubComment, error) {
	var answers []model.SubComment
	err := cd.db.Where("target_id = ?", targetID).Find(&answers).Error
	return answers, err
}
