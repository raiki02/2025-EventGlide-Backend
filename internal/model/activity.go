package model

import "time"

type Activity struct {
	Bid        string    `gorm:"type:varchar(255);primary_key;not null;unique;column:bid"`
	CreatedAt  time.Time `gorm:"type:datetime;column:created_at; not null"`
	IsChecking string    `gorm:"type:varchar(255);column:is_checking"`

	StudentID      string `gorm:"type:varchar(255);column:student_id;not null"`
	Title          string `gorm:"type:varchar(255);column:title;not null"`
	Introduce      string `gorm:"type:text;column:introduce;not null"`
	ShowImg        string `gorm:"type:text;column:show_img"`
	HolderType     string `gorm:"type:varchar(255);column:holder_type;not null"`
	Position       string `gorm:"type:varchar(255);column:position;not null"`
	IfRegister     string `gorm:"type:enum('是','否');column:if_register;not null"`
	RegisterMethod string `gorm:"type:varchar(255);column:register_method"`
	StartTime      string `gorm:"type:datetime;column:start_time;not null"`
	EndTime        string `gorm:"type:datetime;column:end_time;not null"`
	Type           string `gorm:"type:varchar(255);column:type;not null"`
	ActiveForm     string `gorm:"type:varchar(255);column:active_form"`
	Signer         string `gorm:"type:text;column:signer;not null"`

	LikeNum    int `gorm:"type:int;column:like_num;default:0"`
	CollectNum int `gorm:"type:int;column:collect_num;default:0"`
	CommentNum int `gorm:"type:int;column:comment_num;default:0"`
}

type ActivityDraft struct {
	Bid       string    `gorm:"type:varchar(255);primary_key;not null;unique;column:bid"`
	CreatedAt time.Time `gorm:"type:datetime;column:created_at;not null"`

	StudentID      string `gorm:"type:varchar(255);column:student_id"`
	Title          string `gorm:"type:varchar(255);column:title"`
	Introduce      string `gorm:"type:text;column:introduce"`
	ShowImg        string `gorm:"type:text;column:show_img"`
	HolderType     string `gorm:"type:varchar(255);column:holder_type"`
	Position       string `gorm:"type:varchar(255);column:position"`
	IfRegister     string `gorm:"type:enum('yes','no');column:if_register;not null"`
	RegisterMethod string `gorm:"type:varchar(255);column:register_method"`
	StartTime      string `gorm:"type:datetime;column:start_time"`
	EndTime        string `gorm:"type:datetime;column:end_time"`
	Type           string `gorm:"type:varchar(255);column:type"`
	ActiveForm     string `gorm:"type:varchar(255);column:active_form"`
	Signer         string `gorm:"type:text;column:signer"`
}
