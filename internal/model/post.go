package model

type Post struct {
	Id       int    `json:"id" gorm:"column:id; type:int; comment:主键; not null; primaryKey; autoIncrement"`
	Title    string `json:"title" gorm:"column:title; type:varchar(255); comment:标题; not null"`
	Content  string `json:"content" gorm:"column:content; type:text; comment:详细内容; not null"`
	PosterId int    `json:"poster_id" gorm:"column:poster_id; type:int; comment:发布者id; not null"`
}
