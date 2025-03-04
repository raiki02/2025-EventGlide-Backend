package resp

import (
	"github.com/raiki02/EG/internal/model"
	"time"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type LoginResp struct {
	Id     int    `json:"Id"`
	Sid    string `json:"sid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	School string `json:"school"`
	Likes  string `json:"likes"`
	Token  string `json:"token"`
}

type CommentResp struct {
	CreatedAt   time.Time `json:"created_at"`
	CreatorID   string    `json:"creator_id"`
	Content     string    `json:"content"`
	SubComments int       `json:"sub_comments"`
	Likes       int       `json:"likes"`
}

type AnswerResp struct {
	CreatedAt time.Time `json:"created_at"`
	CreatorID string    `json:"creator_id"`
	Likes     int       `json:"likes"`
	Content   string    `json:"content"`
}

type NumberSearchResp struct {
	Total int             `json:"total"`
	Nums  []*model.Number `json:"nums"`
}

type ListActivitiesResp struct {
	User struct {
		Sid      string `json:"sid"`
		Avatar   string `json:"avatar"`
		Username string `json:"username"`
		School   string `json:"school"`
	} `json:"user"`

	DetailTime struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	} `json:"detail_time"`

	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Host        string   `json:"host"`
	Type        string   `json:"type"`
	IfRegister  string   `json:"if_register"`
	ImgUrls     []string `json:"img_urls"`
	Likes       int      `json:"likes"`
	Comments    int      `json:"comments"`
}
