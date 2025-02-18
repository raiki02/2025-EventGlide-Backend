package model

import (
	"time"
)

type Post struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; comment:创建时间"`
	CreatorID string    `json:"creator_id" gorm:"column:creator_id; type:string; comment:创建者id"`
	Bid       string    `json:"bid" gorm:"column:bid; type:string; comment:绑定id"`
	Title     string    `json:"title" gorm:"column:title; type:varchar(255); comment:标题; not null"`
	Content   string    `json:"content" gorm:"column:content; type:text; comment:详细内容; not null"`
	ImgUrls   string    `json:"img_urls" gorm:"column:img_urls; type:text; comment:图片链接"`

	Comments int `json:"comments" gorm:"column:comments; type:int; comment:评论数;default:0"`
	Likes    int `json:"likes" gorm:"column:likes; type:int; comment:点赞数;default:0"`
}

type PostDraft struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; comment:创建时间"`
	CreatorID string    `json:"creator_id" gorm:"column:creator_id; type:string; comment:创建者id"`
	Bid       string    `json:"bid" gorm:"column:bid; type:string; comment:绑定id"`
	Title     string    `json:"title" gorm:"column:title; type:varchar(255); comment:标题; not null"`
	Content   string    `json:"content" gorm:"column:content; type:text; comment:详细内容; not null"`
	//todo db img column delete
}
