package req

type InteractionReq struct {
	TargetID string `json:"targetid" binding:"required"`
	Subject  string `json:"subject"`
}
