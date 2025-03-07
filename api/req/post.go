package req

type CreatePostReq struct {
	StudentID string   `json:"studentid"`
	Title     string   `json:"title"`
	Introduce string   `json:"introduce"`
	ShowImg   []string `json:"showImg"`
}

type FindPostReq struct {
	Name string `json:"name"`
}
