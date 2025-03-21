package model

import "time"

type Feed struct {
	TargetBid string    `gorm:"column:target_bid; type:varchar(255); not null; comment:目标id" json:"target_id"`
	Object    string    `gorm:"column:object; type:varchar(255); not null; comment:目标主题" json:"object"`
	StudentId string    `gorm:"column:student_id; type:varchar(255); not null; comment:学生id" json:"studentid"`
	CreatedAt time.Time `gorm:"column:created_at; type:datetime; not null; comment:创建时间"`
	Action    string    `gorm:"column:action; type:varchar(255); not null; comment:行为"`
	Status    string    `gorm:"column:status; type:varchar(255); not null; comment:状态; default:'未读'"`
}
