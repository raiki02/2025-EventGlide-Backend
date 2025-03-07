package req

// todo 写在api的res/req中

type DraftReq struct {
	Sid string `json:"sid"`
	Bid string `json:"bid"`
}

type UserAvatarReq struct {
	Sid       string `json:"sid"`
	AvatarUrl string `json:"avatar_url"`
}
