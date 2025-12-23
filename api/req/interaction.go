package req

type InteractionReq struct {
	TargetID string `json:"targetid" binding:"required"`
	Receiver string `json:"receiver" binding:"required"`
	Subject  string `json:"subject"`
}
