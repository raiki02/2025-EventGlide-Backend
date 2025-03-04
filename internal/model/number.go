package model

import "time"

type Number struct {
	FromSid string `json:"from_sid" gorm:"column:from_sid; type:varchar(20); not null"`
	ToSid   string `json:"to_sid" gorm:"column:to_sid; type:varchar(20); not null"`

	Object string `json:"object" gorm:"column:object; type:enum('act','post'); not null"`
	Action string `json:"action" gorm:"column:action; type:enum('like','comment'); not null"`

	IsRead    bool      `gorm:"column:is_read; type:tinyint(1); not null; default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; not null; default:CURRENT_TIMESTAMP"`
	Content   string    `json:"content" gorm:"column:content; type:varchar(255); not null"`
}
