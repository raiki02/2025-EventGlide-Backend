package req

type DraftReq struct {
	Sid string `json:"studentid"`
	Bid string `json:"bid"`
}

type UserAvatarReq struct {
	Sid       string `json:"studentid"`
	AvatarUrl string `json:"avatar_url"`
}
