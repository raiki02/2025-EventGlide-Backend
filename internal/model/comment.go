package model

import "time"

type Comment struct {
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at; type:datetime; comment:创建时间; not null"`
	Bid         string    `json:"bid" gorm:"column:bid; type:varchar(255); comment:绑定id ;not null"`
	CreatorID   string    `json:"creator_id" gorm:"column:creator_id; type:varchar(255); comment:创建者id ;not null"`
	TargetID    string    `json:"target_id" gorm:"column:target_id; type:varchar(255); comment:目标id ;not null"`
	Likes       int       `json:"likes" gorm:"column:likes; type:int; comment:点赞数;default:0"`
	SubComments int       `json:"sub_comments" gorm:"column:sub_comments; type:int; comment:回复数;default:0; not null"`
	Content     string    `json:"content" gorm:"column:content; type:text; comment:评论内容; not null"`
}

type SubComment struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; comment:创建时间 ;not null"`
	Bid       string    `json:"bid" gorm:"column:bid; type:varchar(255); comment:绑定id ;not null"`
	CreatorID string    `json:"creator_id" gorm:"column:creator_id; type:varchar(255); comment:创建者id ;not null"`
	TargetID  string    `json:"target_id" gorm:"column:target_id; type:varchar(255); comment:目标id ;not null"`
	Likes     int       `json:"likes" gorm:"column:likes; type:int; comment:点赞数;default:0"`
	Content   string    `json:"content" gorm:"column:content; type:text; comment:回复内容 ;not null"`
}
