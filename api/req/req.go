package req

type DraftReq struct {
	Sid string `json:"studentid"`
	Bid string `json:"bid"`
}

type UserAvatarReq struct {
	StudentID string `json:"studentid"`
	AvatarUrl string `json:"avatar_url"`
}

type AuditWrapper struct {
	Subject  string
	CactReq  *CreateActReq
	CpostReq *CreatePostReq
}
