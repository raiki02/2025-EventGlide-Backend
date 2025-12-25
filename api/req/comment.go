package req

type DeleteCommentReq struct {
	StudentID string `json:"studentid" validate:"required"`
	TargetID  string `json:"target_id" validate:"required"`
}

type CreateCommentReq struct {
	StudentID string `json:"studentid" validate:"required"`
	Content   string `json:"content" validate:"required"`
	ParentID  string `json:"parent_id" validate:"required"`
	Subject   string `json:"subject" validate:"required"`
	Receiver  string `json:"receiver" validate:"required"`
}

type LoadCommentsReq struct {
	Id string `json:"id" validate:"required" form:"id" uri:"id"`
}
