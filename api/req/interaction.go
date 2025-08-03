package req

type InteractionReq struct {
	StudentID string `json:"studentid"`
	TargetID  string `json:"targetid"`
	Subject   string `json:"subject"`
}
