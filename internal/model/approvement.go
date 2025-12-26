package model

import (
	"gorm.io/gorm"
	"log"
	"time"
)

const (
	StancePass   = "pass"
	StanceReject = "reject"
)

type Approvement struct {
	Id  uint   `gorm:"primary_key;auto_increment"`            // 主键ID
	Bid string `gorm:"type:varchar(255);not null;column:bid"` // 活动/帖子ID

	StudentId   string    `gorm:"type:varchar(255);not null;column:student_id"`                                  // 受邀请的学号
	StudentName string    `gorm:"type:varchar(255);not null;column:student_name"`                                // 受邀请人的姓名                    // 学生姓名
	Stance      string    `gorm:"type:enum('pass','reject','pending');default:'pending';column:stance;not null"` // 代表赞同或反对
	UpdatedAt   time.Time `gorm:"type:datetime;column:updated_at;not null"`                                      // 修改代表赞同或反对的时间
	CreatedAt   time.Time `gorm:"type:datetime;column:created_at;not null"`
}

func (a *Approvement) AfterUpdate(tx *gorm.DB) (err error) {
	if a.Stance == StancePass {
		passUpdate := tx.Exec(`
			UPDATE activity
			SET is_checking = 'pass'
			WHERE bid = ?
			AND NOT EXISTS (
				SELECT 1
				FROM approvement
				WHERE bid = ?
				AND stance != 'pass'
			)
			AND EXISTS (
				SELECT 1
				FROM auditor_form
				WHERE bid = ?
				AND status = 'pass'
			)
		`, a.Bid, a.Bid, a.Bid)
		if passUpdate.Error != nil {
			log.Println("approvement AfterUpdate error when passing:", passUpdate.Error)
			return passUpdate.Error
		}
		if passUpdate.RowsAffected > 0 {
			log.Println("approvement AfterUpdate passed successfully for bid:", a.Bid)
			return nil
		}
	} else if a.Stance == StanceReject {
		rejectUpdate := tx.Exec(`
			UPDATE activity
			SET is_checking = 'reject'
			WHERE bid = ?
		`, a.Bid)
		if rejectUpdate.Error != nil {
			log.Println("approvement AfterUpdate error when rejecting:", rejectUpdate.Error)
			return rejectUpdate.Error
		}
		if rejectUpdate.RowsAffected > 0 {
			log.Println("approvement AfterUpdate rejected successfully for bid:", a.Bid)
			return nil
		}
	}
	return nil
}
