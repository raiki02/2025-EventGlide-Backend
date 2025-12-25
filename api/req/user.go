package req

type UpdateNameReq struct {
	Name string `json:"new_name"`
}
type LoginReq struct {
	StudentID string `json:"studentid" validate:"required,len=10"`
	Password  string `json:"password" validate:"required"`
}

type UserSearchReq struct {
	Keyword string `json:"keyword"`
}
