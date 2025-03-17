package resp

type CreateActivityResp struct {
	Bid        string   `json:"bid"`
	Title      string   `json:"title"`
	Introduce  string   `json:"introduce"`
	ShowImg    []string `json:"showImg"`
	Type       string   `json:"type"`
	Position   string   `json:"position"`
	IfRegister string   `json:"if_register"`
	IsChecking string   `json:"isChecking"`
	ActiveForm string   `json:"activeForm"`
	Signer     []struct {
		StudentID string `json:"studentid"`
		Name      string `json:"name"`
	} `json:"signer"`
	UserInfo struct {
		StudentID string `json:"studentid"`
		Avatar    string `json:"avatar"`
		Username  string `json:"username"`
		School    string `json:"school"`
	} `json:"userInfo"`
}

type ListActivitiesResp struct {
	UserInfo struct {
		StudentID string `json:"studentid"`
		Avatar    string `json:"avatar"`
		Username  string `json:"username"`
		School    string `json:"school"`
	} `json:"userInfo"`

	DetailTime struct {
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
	} `json:"detailTime"`

	Title      string   `json:"title"`
	Bid        string   `json:"bid"`
	Introduce  string   `json:"introduce"`
	Position   string   `json:"position"`
	Type       string   `json:"type"`
	HolderType string   `json:"holderType"`
	IfRegister string   `json:"if_register"`
	ShowImg    []string `json:"showImg"`

	LikeNum    int `json:"likeNum"`
	CollectNum int `json:"collectNum"`
	CommentNum int `json:"commentNum"`

	IsLike    string `json:"isLike"`
	IsCollect string `json:"isCollect"`
}
