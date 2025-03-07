package req

type UpdateNameReq struct {
	StudentID string `json:"studentid"`
	Name      string `json:"new_name"`
}
type LoginReq struct {
	StudentiD string `json:"studentid"`
	Password  string `json:"password"`
}

type UserSearchReq struct {
	StudentiD string `json:"studentid"`
	Keyword   string `json:"keyword"`
}
