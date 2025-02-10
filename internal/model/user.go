package model

// 每个首次使用的用户根据插入数据库，用于活动草稿，头像，昵称操作
type User struct {
	Id        int    `gorm:"column:id; type: int; not null; primary_key; autoIncrement" json:"id"`
	StudentId string `gorm:"column:sid;type:varchar(255);not null" json:"sid"`
	Name      string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`
	School    string `gorm:"column:school;type:varchar(255);not null" json:"school"`
}

const (
	DefaultAvatar string = "http://img.inside-me.top/EventGlide/OIP-C.jpg"
)
