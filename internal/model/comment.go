package model

import (
	"gorm.io/gorm"
	"time"
)

/*
	post -> comment -> subcomment 独立

在一个帖子下评论，获取帖子id，制作评论发布：评论id，发布者id，评论内容
回复评论，获取帖子id，评论id，制作回复：回复id，回复者id，回复内容

act-post -> comment -> subcomment 共享
获取bid，制作评论发布：评论id，发布者id，评论内容
回复评论，获取bid，评论id，制作回复：回复id，回复者id，回复内容
*/
type Comment struct {
	CreatedAt time.Time `json:"created_at" gorm:"column: created_at; not null; type: datetime; comment: 创建时间"`
	DeletedAt gorm.DeletedAt

	CommentID string `json:"comment_id" gorm:"column: comment_id; not null; type: varchar(255); comment: 评论ID;primary_key"`
	PosterID  string `json:"poster_id" gorm:"column: poster_id; not null; type: varchar(255); comment: 发布者ID"`
	TargetID  string `json:"target_id" gorm:"column: target_id; not null; type: varchar(255); comment: 帖子id或者绑定id"`
	ParentID  string `json:"parent_id" gorm:"column: parent_id; type: varchar(255); comment: 父评论id,专用于回复"`

	Content string `json:"content" gorm:"column: content; not null; type: text; comment: 评论内容"`

	Likes   int `json:"likes" gorm:"column: likes; default:0; type: int; comment: 点赞数"`
	Answers int `json:"answers" gorm:"column: answers; default:0; type: int; comment: 回复数"`
}
