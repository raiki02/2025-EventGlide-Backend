package req

type NumReq struct {
	StudentID string `json:"studentid"`
	TargetId  string `json:"target_id"`
	Object    string `json:"object"`
	Action    string `json:"action"`
}
