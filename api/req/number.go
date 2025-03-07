package req

type NumberSendReq struct {
	FromSid string `json:"from_sid"`
	ToSid   string `json:"to_sid"`
	Object  string `json:"object"`
	Action  string `json:"action"`
	Content string `json:"content"`
}

type NumberSearchReq struct {
	Sid    string `json:"sid"`
	Object string `json:"object"`
	Action string `json:"action"`
}

type NumberDelReq struct {
	Sid    string `json:"sid"`
	Object string `json:"object"`
}

type NumReq struct {
	TargetId string `json:"target_id"`
	Object   string `json:"object"`
}
