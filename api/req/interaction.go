package req

type InteractionReq struct {
	TargetID string `json:"targetid" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	Subject  string `json:"subject" validate:"required"`
}
