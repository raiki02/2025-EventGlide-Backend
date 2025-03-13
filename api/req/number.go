package req

type NumberSendReq struct {
	FromSid string `json:"from_sid"`
	ToSid   string `json:"to_sid"`
	Object  string `json:"object"`
	Action  string `json:"action"`
	Content string `json:"content"`
}

type NumberSearchReq struct {
	StudentID string `json:"studentid"`
	Object    string `json:"object"`
	Action    string `json:"action"`
}

type NumberDelReq struct {
	StudentID string `json:"studentid"`
	Object    string `json:"object"`
}

type NumReq struct {
	StudentID string `json:"studentid"`
	TargetId  string `json:"target_id"`
	Object    string `json:"object"`
	Type      int    `json:"type"`
}
