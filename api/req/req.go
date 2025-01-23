package req

// todo 写在api的res/req中
type EditUsername struct {
	StudentId string `json:"student_id"`
	NewName   string `json:"new_name"`
}

type EditAvatar struct {
	StudentId string `json:"student_id"`
	NewAvatar string `json:"new_avatar"`
}
