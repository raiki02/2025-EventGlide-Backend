package resp

type CommentResp struct {
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
}
