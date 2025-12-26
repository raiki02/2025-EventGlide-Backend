package req

type NumReq struct {
	TargetId string `json:"target_id"`
	Object   string `json:"object"`
	Action   string `json:"action"`
}

type ReadFeedDetailReq struct {
	Id string `json:"id" validate:"required" form:"id" uri:"id"`
}
