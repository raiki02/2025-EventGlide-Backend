package req

type DraftReq struct {
	Bid string `json:"bid"`
}

type UserAvatarReq struct {
	AvatarUrl string `json:"avatar_url"`
}

type AuditWrapper struct {
	Subject  string
	CactReq  *CreateActReq
	CpostReq *CreatePostReq
}
