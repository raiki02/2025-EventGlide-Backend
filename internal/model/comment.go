package model

import "time"

/*
	post -> comment -> subcomment 独立

在一个帖子下评论，获取帖子id，制作评论发布：评论id，发布者id，评论内容
回复评论，获取帖子id，评论id，制作回复：回复id，回复者id，回复内容

act-post -> comment -> subcomment 共享
获取bid，制作评论发布：评论id，发布者id，评论内容
回复评论，获取bid，评论id，制作回复：回复id，回复者id，回复内容
*/
//for post/act
type Comment struct {
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at; type:datetime; comment:创建时间; not null"`
	Bid         string    `json:"bid" gorm:"column:bid; type:varchar(255); comment:绑定id ;not null"`
	CreatorID   string    `json:"creator_id" gorm:"column:creator_id; type:varchar(255); comment:创建者id ;not null"`
	TargetID    string    `json:"target_id" gorm:"column:target_id; type:varchar(255); comment:目标id ;not null"`
	Likes       int       `json:"likes" gorm:"column:likes; type:int; comment:点赞数;default:0"`
	SubComments int       `json:"sub_comments" gorm:"column:sub_comments; type:int; comment:回复数;default:0; not null"`
	Content     string    `json:"content" gorm:"column:content; type:text; comment:评论内容; not null"`
}

// for comment
type SubComment struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; comment:创建时间 ;not null"`
	Bid       string    `json:"bid" gorm:"column:bid; type:varchar(255); comment:绑定id ;not null"`
	CreatorID string    `json:"creator_id" gorm:"column:creator_id; type:varchar(255); comment:创建者id ;not null"`
	TargetID  string    `json:"target_id" gorm:"column:target_id; type:varchar(255); comment:目标id ;not null"`
	Likes     int       `json:"likes" gorm:"column:likes; type:int; comment:点赞数;default:0"`
	Content   string    `json:"content" gorm:"column:content; type:text; comment:回复内容 ;not null"`
}
