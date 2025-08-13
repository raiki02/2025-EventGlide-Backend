package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	SubjectActivity = "activities"
	SubjectPost     = "posts"
)

type AuditorForm struct {
	Id        uint      `gorm:"primary_key;auto_increment;not null"`
	Subject   string    `gorm:"type:varchar(255);not null"`                                                       // 活动 or 帖子
	Bid       string    `gorm:"type:varchar(255);not null;unique;column:bid"`                                     // 活动/帖子ID
	Status    string    `gorm:"type:enum('pending','approve','reject');default:'pending';column:status;not null"` // 表单审核状态 审核是0,1,2
	FormUrl   string    `gorm:"type:text;column:form_url"`                                                        // 表单的URL地址 // 给活动用的填报表单
	CreatedAt time.Time `gorm:"type:datetime;column:created_at;not null"`                                         // 创建时间
	UpdatedAt time.Time `gorm:"type:datetime;column:updated_at;not null"`                                         // 更新时间
}

func (af *AuditorForm) AfterUpdate(tx *gorm.DB) (err error) {
	if af.Status == "approve" {
		table := af.Subject
		if table == SubjectActivity {
			update := tx.Exec(`
				UPDATE activities
				SET is_checking = 'pass'
				WHERE bid = ?
				AND NOT EXISTS (
					SELECT 1
					FROM approvements
					WHERE bid = ?
					AND stance != 'approve'
				)
			`, af.Bid, af.Bid)
			if update.Error != nil {
				return update.Error
			}
			if update.RowsAffected > 0 {
				return nil
			}
		} else if table == SubjectPost {
			update := tx.Exec(`
				UPDATE posts
				SET is_checking = 'pass'
				WHERE bid = ?
			`, af.Bid)
			if update.Error != nil {
				return update.Error
			}
			if update.RowsAffected > 0 {
				return nil
			}
		}
	}
	return nil
}
