package model

import "gorm.io/gorm"

type Post struct {
	*gorm.Model
	BId        int    `json:"bid" gorm:"column:bid; type:int; comment:绑定id; not null; primaryKey; autoIncrement"`
	Title      string `json:"title" gorm:"column:title; type:varchar(255); comment:标题; not null"`
	Content    string `json:"content" gorm:"column:content; type:text; comment:详细内容; not null"`
	PosterId   int    `json:"poster_id" gorm:"column:poster_id; type:int; comment:发布者id; not null"`
	ImgUrls    string `json:"img_urls" gorm:"column:img_urls; type:text; comment:图片链接; not null"`
	CommentNum int    `json:"comment_num" gorm:"column:comment_num; type:int; comment:评论数; not null"`
	LikeNum    int    `json:"like_num" gorm:"column:like_num; type:int; comment:点赞数; not null"`
}
