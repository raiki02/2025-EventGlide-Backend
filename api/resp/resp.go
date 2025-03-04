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
