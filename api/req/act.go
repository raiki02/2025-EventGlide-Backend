package req

type ActSearchReq struct {
	Type       []string `json:"type"`
	Host       []string `json:"host"`
	Location   []string `json:"location"`
	IfRegister string   `json:"if_register"`
	DetailTime struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	} `json:"detailTime"`
}

type CreateActReq struct {
	StudentID string   `json:"studentid"`
	Title     string   `json:"title"`
	Introduce string   `json:"introduce"`
	ShowImg   []string `json:"showImg"`

	LabelForm struct {
		HolderType     string `json:"hodlerType"`
		Position       string `json:"position"`
		IfRegister     string `json:"if_register"`
		RegisterMethod string `json:"register_method"`
		StartTime      string `json:"startTime"`
		ActiveForm     string `json:"activeForm"`
		EndTime        string `json:"endTime"`
		Type           string `json:"type"`

		Signer []struct {
			StudentID string `json:"studentid"`
			Name      string `json:"name"`
		} `json:"signer"`
	} `json:"labelform"`
}
