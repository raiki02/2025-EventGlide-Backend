package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
)

type ActivityServiceHdl interface {
	NewAct()
	NewDraft()

	FindActBySearches(*gin.Context, *req.ActSearchReq) ([]model.Activity, error)
	ShowActDetails(*gin.Context, string) (model.Activity, error)
}

type ActivityService struct {
	ad *dao.ActDao
}

func NewActivityService(ad *dao.ActDao) *ActivityService {
	return &ActivityService{
		ad: ad,
	}
}

func (as *ActivityService) FindActBySearches(ctx *gin.Context, req *req.ActSearchReq) ([]model.Activity, error) {
	temp := &model.Activity{
		Type:       req.Type,
		StartTime:  req.StartTime + ":00", // 2021-06-01 00:00:00
		EndTime:    req.EndTime + ":00",
		Host:       req.Host,
		Location:   req.Location,
		IfRegister: req.IfRegister,
	}
	acts, err := as.ad.FindActBySearches(ctx, temp)
	if err != nil {
		return nil, err
	}
	return acts, nil
}

func (as *ActivityService) ShowActDetails(c *gin.Context, bid string) (model.Activity, error) {
	act, err := as.ad.FindActByBid(c, bid)
	if err != nil {
		return model.Activity{}, err
	}
	return act, nil

}
