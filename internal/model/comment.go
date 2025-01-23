package model

import (
	"time"
)

type Comment struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; comment:创建时间"`
	Content   string    `json:"content" gorm:"column:content; type:text; comment:详细内容"`
	PosterId  int       `json:"poster_id" gorm:"column:poster_id; type:int; comment:发布者id"`
	PostId    int       `json:"post_id" gorm:"column:post_id; type:int; comment:帖子id"`
}
