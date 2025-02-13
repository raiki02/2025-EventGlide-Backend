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
