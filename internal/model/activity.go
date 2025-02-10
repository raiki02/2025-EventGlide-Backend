package model

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	// 活动id
	*gorm.Model
	CreatorId string `json:"creator_id" gorm:"not null;type: varchar(255);comment:创建者id;column:creator_id"`
	Bid       string `json:"bid" gorm:"not null;type: varchar(255);comment:绑定id;column:bid"`
	//divided by function
	//basic
	Type           string `json:"type" gorm:"not null;type: varchar(255);comment:活动类型;column:type"`
	Host           string `json:"host" gorm:"not null;type: varchar(255);comment:活动主办方;column:host"`
	Location       string `json:"location" gorm:"not null;type: varchar(255);comment:活动地点;column:location"`
	IfRegister     string `json:"if_register" gorm:"null;type: enum('yes','no');comment:是否需要报名;column:if_register"`
	RegisterMethod string `json:"register_method" gorm:"null;type: varchar(255);comment:报名方式;column:register_method"`

	//complex
	StartTime string `json:"start_time" gorm:"not null;type: datetime;comment:活动开始时间;column:start_time"`
	EndTime   string `json:"end_time" gorm:"not null;type: datetime;comment:活动结束时间;column:end_time"`
	Name      string `json:"name" gorm:"not null;type: varchar(255);comment:活动名称;column:name"`

	//descriptive
	Capacity int    `json:"capacity" gorm:"not null;type: int;comment:活动容量;column:capacity"`
	Images   string `json:"images" gorm:"null; type: text; comment:图片描述; column:images“`
	Details  string `json:"details" gorm:"not null;type: text;comment:活动详情;column:details"`

	//interactive
	Likes    int `json:"likes" gorm:"default:0;type: int;comment:活动点赞数;column:likes"`
	Comments int `json:"comments" gorm:"default:0;type: int;comment:活动评论数;column:comments"`

	//audit
	Identification string `json:"identification"`
	Audition       string `json:"audition"`
}

func (act Activity) SetBid(ctx *gin.Context) error {
	bid, err := uuid.NewUUID()
	if err != nil {
		act.Bid = ""
		return err
	}
	act.Bid = bid.String()
	return nil
}

func (act Activity) GetImgUrl(ctx *gin.Context) error {
	return nil
}

type ActivityDraft struct {
	*gorm.Model
	Images         string `json:"images" gorm:"null; type: text; comment:图片描述; column:images“`
	Name           string `json:"name" gorm:"not null;type: varchar(255);comment:活动名称;column:name"`
	Details        string `json:"details" gorm:"not null;type: text;comment:活动详情;column:details"`
	Type           string `json:"type" gorm:"not null;type: varchar(255);comment:活动类型;column:type"`
	Host           string `json:"host" gorm:"not null;type: varchar(255);comment:活动主办方;column:host"`
	Location       string `json:"location" gorm:"not null;type: varchar(255);comment:活动地点;column:location"`
	IfRegister     string `json:"if_register" gorm:"null;type: enum('yes','no');comment:是否需要报名;column:if_register"`
	RegisterMethod string `json:"register_method" gorm:"null;type: varchar(255);comment:报名方式;column:register_method"`
	StartTime      string `json:"start_time" gorm:"not null;type: datetime;comment:活动开始时间;column:start_time"`
	EndTime        string `json:"end_time" gorm:"not null;type: datetime;comment:活动结束时间;column:end_time"`
	Capacity       int    `json:"capacity" gorm:"not null;type: int;comment:活动容量;column:capacity"`
	Visibility     string `json:"visibility" gorm:"not null;type: varchar(255);comment:活动可见性;column:visibility"`
	Likes          int    `json:"likes" gorm:"not null;type: int;comment:活动点赞数;column:likes"`
	Comments       int    `json:"comments" gorm:"not null;type: int;comment:活动评论数;column:comments"`

	Identification string `json:"identification"`
	Audition       string `json:"audition"`
}
