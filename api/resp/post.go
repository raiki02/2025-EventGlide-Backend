package resp

type CreatePostResp struct {
	Bid         string `json:"bid"`
	StudentID   string `json:"studentid"`
	PublishTime string `json:"publishTime"`

	Title      string   `json:"title"`
	Introduce  string   `json:"introduce"`
	ShowImg    []string `json:"showImg"`
	IsChecking string   `json:"isChecking"`

	UserInfo struct {
		StudentID string `json:"studentid"`
		Avatar    string `json:"avatar"`
		Username  string `json:"username"`
		School    string `json:"school"`
	} `json:"userInfo"`
}

type ListPostsResp struct {
	Bid      string `json:"bid"`
	UserInfo struct {
		StudentID string `json:"studentid"`
		Avatar    string `json:"avatar"`
		Username  string `json:"username"`
		School    string `json:"school"`
	} `json:"userInfo"`
	PublishTime string `json:"publishTime"`

	Introduce string   `json:"introduce"`
	ShowImg   []string `json:"showImg"`
	Title     string   `json:"title"`

	LikeNum    int `json:"likeNum"`
	CollectNum int `json:"collectNum"`
	CommentNum int `json:"commentNum"`

	IsLike     string `json:"isLike"`
	IsCollect  string `json:"isCollect"`
	IsChecking string `json:"isChecking"`
}
