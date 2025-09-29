package req

type UpdateNameReq struct {
	Name string `json:"new_name"`
}
type LoginReq struct {
	StudentID string `json:"studentid"`
	Password  string `json:"password"`
}

type UserSearchReq struct {
	Keyword string `json:"keyword"`
}
