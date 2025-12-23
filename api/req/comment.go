package req

type DeleteCommentReq struct {
	StudentID string `json:"studentid"`
	TargetID  string `json:"target_id"`
}

type CreateCommentReq struct {
	StudentID string `json:"studentid"`
	Content   string `json:"content"`
	ParentID  string `json:"parent_id"`
	Subject   string `json:"subject"`
	Receiver  string `json:"receiver"`
}
