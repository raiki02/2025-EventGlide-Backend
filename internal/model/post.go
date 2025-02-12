package model

import (
	"time"
)

type Post struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; comment:创建时间"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at; type:datetime; comment:删除时间"`

	BId      int    `json:"bid" gorm:"column:bid; type:int; comment:绑定id"`
	Title    string `json:"title" gorm:"column:title; type:varchar(255); comment:标题; not null"`
	Content  string `json:"content" gorm:"column:content; type:text; comment:详细内容; not null"`
	PosterId string `json:"poster_id" gorm:"column:poster_id; type:varchar(255); comment:发布者id; not null"`
	ImgUrls  string `json:"img_urls" gorm:"column:img_urls; type:text; comment:图片链接; not null"`

	Comments int `json:"comments" gorm:"column:comments; type:int; comment:评论数;default:0"`
	Likes    int `json:"likes" gorm:"column:likes; type:int; comment:点赞数;default:0"`
}
