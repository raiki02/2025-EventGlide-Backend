package req

// todo 写在api的res/req中

type NumberReq struct {
	Topic     string `json:"topic"`
	Msg       string `json:"msg"`
	ExcuterID string `json:"excuter_id"`
}

type ActSearchReq struct {
	Type       string `json:"type"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Host       string `json:"host"`
	Location   string `json:"location"`
	IfRegister string `json:"if_register"`
}

func (asr *ActSearchReq) ToMap() map[string]string {
	m := make(map[string]string)
	m["type"] = asr.Type
	m["start_time"] = asr.StartTime
	m["end_time"] = asr.EndTime
	m["host"] = asr.Host
	m["location"] = asr.Location
	m["if_register"] = asr.IfRegister
	return m
}

type UserSearchReq struct {
	Sid     string `json:"sid"`
	Keyword string `json:"keyword"`
}

type DraftReq struct {
	Sid string `json:"sid"`
	Bid string `json:"bid"`
}

type CommentReq struct {
	CreatorID string `json:"creator_id"`
	TargetID  string `json:"target_id"`
	Content   string `json:"content"`
}
