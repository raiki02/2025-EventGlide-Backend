package resp

import "github.com/raiki02/EG/internal/model"

type NumberSearchResp struct {
	Total int             `json:"total"`
	Nums  []*model.Number `json:"nums"`
}
