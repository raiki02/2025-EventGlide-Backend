package resp

type BriefFeedResp struct {
	LikeAndCollect int `json:"likeandcollect"`
	CommentAndAt   int `json:"commentandat"`
	Total          int `json:"total"`
}

type FeedResp struct {
	Likes    []FeedLikeResp    `json:"likes"`
	Ats      []FeedAtResp      `json:"ats"`
	Comments []FeedCommentResp `json:"comments"`
	Collects []FeedCollectResp `json:"collects"`
}

type UserInfo struct {
	StudentID string `json:"studentid"`
	Avatar    string `json:"avatar"`
	Username  string `json:"username"`
}
type FeedLikeResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	TargetBid   string `json:"target_bid"`
	Status      string `json:"status"`
}

type FeedCommentResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	TargetBid   string `json:"target_bid"`
	Status      string `json:"status"`
}

type FeedAtResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	TargetBid   string `json:"target_bid"`
	Status      string `json:"status"`
}

type FeedCollectResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	TargetBid   string `json:"target_bid"`
	Status      string `json:"status"`
}
