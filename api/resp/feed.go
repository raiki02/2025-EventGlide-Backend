package resp

type BriefFeedResp struct {
	LikeAndCollect int `json:"likeandcollect"`
	CommentAndAt   int `json:"commentandat"`
	Total          int `json:"total"`
}

type FeedResp struct {
	Likes       []FeedLikeResp       `json:"likes,omitempty"`
	Ats         []FeedAtResp         `json:"ats,omitempty"`
	Comments    []FeedCommentResp    `json:"comments,omitempty"`
	Collects    []FeedCollectResp    `json:"collects,omitempty"`
	Invitations []FeedInvitationResp `json:"invitations,omitempty"`
}

type UserInfo struct {
	StudentID string `json:"studentid"`
	Avatar    string `json:"avatar"`
	Username  string `json:"username"`
}
type FeedLikeResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Id          int64  `json:"id"`
	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	TargetBid   string `json:"target_bid"`
	FirstPic    string `json:"first_pic,omitempty"`
	Status      string `json:"status"`
}

type FeedCommentResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Id          int64  `json:"id"`
	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	TargetBid   string `json:"target_bid"`
	FirstPic    string `json:"first_pic,omitempty"`
	Status      string `json:"status"`
}

type FeedAtResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Id          int64  `json:"id"`
	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	TargetBid   string `json:"target_bid"`
	FirstPic    string `json:"first_pic,omitempty"`
	Status      string `json:"status"`
}

type FeedCollectResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Id          int64  `json:"id"`
	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	FirstPic    string `json:"first_pic,omitempty"`
	TargetBid   string `json:"target_bid"`
	Status      string `json:"status"`
}

type FeedInvitationResp struct {
	Userinfo UserInfo `json:"userInfo"`

	Id          int64  `json:"id"`
	Message     string `json:"message"`
	PubLishedAt string `json:"published_at"`
	TargetBid   string `json:"target_bid"`
	FirstPic    string `json:"first_pic,omitempty"`
	Status      string `json:"status"`
}
