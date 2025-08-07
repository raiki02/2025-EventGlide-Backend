package model

import "time"

type Approvement struct {
	Id          uint      `gorm:"primary_key;auto_increment"`
	ActivityId  string    `gorm:"type:varchar(255);not null;unique;column:activity_id"`
	StudentId   string    `gorm:"type:varchar(255);not null;column:student_id"`
	StudentName string    `gorm:"type:varchar(255);not null;column:student_name"` // 学生姓名
	Stance      string    `gorm:"type:enum('赞同','反对');column:stance;not null"`    // 赞同或反对
	UpdatedAt   time.Time `gorm:"type:datetime;column:updated_at;not null"`       // 修改代表赞同或反对的时间
	CreatedAt   time.Time `gorm:"type:datetime;column:created_at;not null"`
}
