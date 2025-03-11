package model

type User struct {
	Id          int    `gorm:"column:id; type: int; not null; primary_key; autoIncrement"`
	StudentID   string `gorm:"column:student_id;type:varchar(255);not null" json:"student_id"`
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Avatar      string `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`
	School      string `gorm:"column:school;type:varchar(255);not null" json:"school"`
	CollectAct  string `gorm:"column:collect_act;type:text" json:"collectact"`
	CollectPost string `gorm:"column:collect_post;type:text" json:"collectpost"`
}
