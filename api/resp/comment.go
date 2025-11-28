package resp

type CommentResp struct {
	Bid string `json:"bid"`

	Creator struct {
		StudentID string `json:"studentid"`
		Username  string `json:"username"`
		Avatar    string `json:"avatar"`
	} `json:"creator"`

	CommentedTime string `json:"commented_time"`
	CommentedPos  string `json:"commented_pos"`
	Content       string `json:"content"`
	LikeNum       int    `json:"likeNum"`
	ReplyNum      int    `json:"replyNum"`
	IsLike        string `json:"isLike"`
	ParentID      string `json:"parent_id"`
	RootID        string `json:"root_id"`

	Reply []ReplyResp `json:"reply"`
}

type ReplyResp struct {
	Bid string `json:"bid"`

	ReplyCreator struct {
		StudentID string `json:"studentid"`
		Username  string `json:"username"`
		Avatar    string `json:"avatar"`
	} `json:"reply_creator"`

	ReplyContent string `json:"reply_content"`
	ReplyTime    string `json:"reply_time"`
	ReplyPos     string `json:"reply_pos"`

	ParentID       string `json:"parent_id"`
	RootID         string `json:"root_id"`
	ParentUserName string `json:"parentUserName"`

	IsLike   string `json:"isLike"`
	LikeNum  int    `json:"likeNum"`
	ReplyNum int    `json:"replyNum"`
}
