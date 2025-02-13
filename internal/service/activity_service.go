package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/api/resp"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
)

type ActivityServiceHdl interface {
	NewAct()
	NewDraft()

	FindActBySearches(*gin.Context, *req.ActSearchReq) ([]model.Activity, error)
	ShowActDetails(*gin.Context, string) (model.Activity, error)
	FindActByDate(*gin.Context, string) ([]resp.BriefAct, error)
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

func (as *ActivityService) FindActByDate(c *gin.Context, date string) ([]resp.BriefAct, error) {
	acts, err := as.ad.FindActByDate(c, date)
	if err != nil {
		return nil, err
	}
	return daoTOresp(acts), nil
}

func daoTOresp(acts []model.Activity) []resp.BriefAct {
	var resps []resp.BriefAct
	for _, act := range acts {
		resp := resp.BriefAct{
			Bid:       act.Bid,
			Location:  act.Location,
			StartTime: act.StartTime,
			EndTime:   act.EndTime,
			Title:     act.Name,
		}
		resps = append(resps, resp)
	}
	return resps
}
