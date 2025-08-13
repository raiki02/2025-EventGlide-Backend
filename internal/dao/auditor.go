package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type AuditorRepository interface {
	Insert(c *gin.Context, bid string, formUrl string, sub string) (*model.AuditorForm, error)
	Update(c *gin.Context, formId uint, status string) error
	Get(c *gin.Context, bid string) (model.AuditorForm, error)
	IsRejected(c *gin.Context, bid string) (bool, error)
}
type AuditorRepo struct {
	db *gorm.DB
}

func NewAuditorRepo(db *gorm.DB) AuditorRepository {
	return &AuditorRepo{
		db: db,
	}
}

func (a *AuditorRepo) Insert(c *gin.Context, bid string, formUrl string, sub string) (*model.AuditorForm, error) {
	form := model.AuditorForm{
		Bid:     bid,
		FormUrl: formUrl,
		Subject: sub,
	}
	if err := a.db.WithContext(c).Create(&form).Error; err != nil {
		return nil, err
	}

	return &form, nil
}

func (a *AuditorRepo) Update(c *gin.Context, formId uint, status string) error {
	return a.db.WithContext(c).Model(&model.AuditorForm{}).Where("id = ?", formId).Update("status", status).Error
}

func (a *AuditorRepo) Get(c *gin.Context, bid string) (model.AuditorForm, error) {
	var form model.AuditorForm
	err := a.db.WithContext(c).Where("bid = ?", bid).First(&form).Error
	return form, err
}

func (a *AuditorRepo) IsRejected(c *gin.Context, bid string) (bool, error) {
	var form model.AuditorForm
	err := a.db.WithContext(c).Where("bid = ? and status = ?", bid, "rejected").First(&form).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // Not rejected
	}
	return true, err // Either found or another error occurred
}
