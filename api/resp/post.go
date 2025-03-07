package resp

type CreatePostResp struct {
	Bid       string   `json:"bid"`
	StudentID string   `json:"studentid"`
	Title     string   `json:"title"`
	Introduce string   `json:"introduce"`
	ShowImg   []string `json:"showImg"`

	UserInfo struct {
		StudentID string `json:"studentid"`
		Avatar    string `json:"avatar"`
		Username  string `json:"username"`
		School    string `json:"school"`
	} `json:"userInfo"`
}

type ListPostsResp struct {
	UserInfo struct {
		StudentID string `json:"studentid"`
		Avatar    string `json:"avatar"`
		Username  string `json:"username"`
		School    string `json:"school"`
	} `json:"userInfo"`

	Introduce string   `json:"introduce"`
	ShowImg   []string `json:"showImg"`
	Title     string   `json:"title"`

	LikeNum    int `json:"likeNum"`
	CollectNum int `json:"collectNum"`
	CommentNum int `json:"commentNum"`
}
