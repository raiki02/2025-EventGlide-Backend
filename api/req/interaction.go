package req

type InteractionReq struct {
	StudentID string `json:"studentid"`
	TargetID  string `json:"targetid" binding:"required"`
	Subject   string `json:"subject"`
}
