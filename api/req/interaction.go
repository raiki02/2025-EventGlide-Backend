package req

type InteractionReq struct {
	TargetID string `json:"targetid" validate:"required"`
	Subject  string `json:"subject" validate:"required"`
}
