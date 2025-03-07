package resp

type LoginResp struct {
	Id      int      `json:"Id"`
	Sid     string   `json:"sid"`
	Name    string   `json:"name"`
	Avatar  string   `json:"avatar"`
	School  string   `json:"school"`
	Collect []string `json:"collect"`
	Token   string   `json:"token"`
}
