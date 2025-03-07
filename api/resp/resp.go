package resp

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ImgBedResp struct {
	DomainName  string `json:"domain_name"`
	AccessToken string `json:"access_token"`
}
