package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/api/resp"
	"github.com/raiki02/EG/internal/cache"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
	"strings"
	"time"
)

type ActivityServiceHdl interface {
	NewAct(*gin.Context, *model.Activity) error
	NewDraft(*gin.Context, *model.ActivityDraft) error
	LoadDraft(*gin.Context, req.DraftReq) (*model.ActivityDraft, error)
	FindActBySearches(*gin.Context, *req.ActSearchReq) ([]resp.ListActivitiesResp, error)
	FindActByDate(*gin.Context, string) ([]resp.ListActivitiesResp, error)
	FindActByName(*gin.Context, string) ([]resp.ListActivitiesResp, error)
	FindActByOwnerID(*gin.Context, string) ([]resp.ListActivitiesResp, error)
	ListAllActs(*gin.Context) ([]resp.ListActivitiesResp, error)
	ToResp(*gin.Context, []model.Activity) []resp.ListActivitiesResp
}

type ActivityService struct {
	ad *dao.ActDao
	ch *cache.Cache
	ud *dao.UserDao
}

func NewActivityService(ad *dao.ActDao, ch *cache.Cache, ud *dao.UserDao) *ActivityService {
	return &ActivityService{
		ad: ad,
		ch: ch,
		ud: ud,
	}
}

func (as *ActivityService) NewAct(c *gin.Context, act *model.Activity) error {
	act.Bid = tools.GenUUID(c)

	createdAt := time.Now()
	act.CreatedAt = createdAt
	err := as.ad.CreateAct(c, act)
	if err != nil {
		return err
	}
	return nil

}

func (as *ActivityService) NewDraft(c *gin.Context, d *model.ActivityDraft) error {
	d.Bid = tools.GenUUID(c)
	err := as.ad.CreateDraft(c, d)
	if err != nil {
		return err
	}
	return nil
}

func (as *ActivityService) LoadDraft(c *gin.Context, dReq req.DraftReq) (model.ActivityDraft, error) {
	d, err := as.ad.LoadDraft(c, dReq.Sid, dReq.Bid)
	if err != nil {
		return model.ActivityDraft{}, err
	}
	return d, nil
}

func (as *ActivityService) FindActBySearches(c *gin.Context, req *req.ActSearchReq) ([]resp.ListActivitiesResp, error) {
	acts, err := as.ad.FindActBySearches(c, req)
	if err != nil {
		return nil, err
	}
	res := as.ToResp(c, acts)
	return res, nil
}

func (as *ActivityService) FindActByDate(c *gin.Context, date string) ([]resp.ListActivitiesResp, error) {
	acts, err := as.ad.FindActByDate(c, date)
	if err != nil {
		return nil, err
	}
	res := as.ToResp(c, acts)
	return res, nil
}

func (as *ActivityService) FindActByName(c *gin.Context, name string) ([]resp.ListActivitiesResp, error) {
	acts, err := as.ad.FindActByName(c, name)
	if err != nil {
		return nil, err
	}
	res := as.ToResp(c, acts)
	return res, nil
}

func (as *ActivityService) FindActByOwnerID(c *gin.Context, id string) ([]resp.ListActivitiesResp, error) {
	acts, err := as.ad.FindActByOwnerID(c, id)
	if err != nil {
		return nil, err
	}
	res := as.ToResp(c, acts)
	return res, nil
}

func (as *ActivityService) ListAllActs(c *gin.Context) ([]resp.ListActivitiesResp, error) {
	acts, err := as.ad.ListAllActs(c)
	if err == nil {
		return nil, err
	}
	res := as.ToResp(c, acts)
	return res, nil
}

func processUrls(c *gin.Context, urls string) []string {
	var res []string
	if urls == "" {
		return res
	}
	return strings.Split(urls, ",")
}

func (as *ActivityService) ToResp(c *gin.Context, acts []model.Activity) []resp.ListActivitiesResp {
	var res []resp.ListActivitiesResp
	for n, act := range acts {
		user := as.ud.FindUserByID(c, act.CreatorId)
		res[n].User.School = user.School
		res[n].User.Username = user.Name
		res[n].User.Avatar = user.Avatar
		res[n].User.Sid = user.StudentId
		res[n].DetailTime.StartTime = act.StartTime
		res[n].DetailTime.EndTime = act.EndTime
		res[n].Host = act.Host
		res[n].Title = act.Name
		res[n].Description = act.Description
		res[n].Location = act.Location
		res[n].Type = act.Type
		res[n].Likes = act.Likes
		res[n].Comments = act.Comments
		res[n].IfRegister = act.IfRegister
		res[n].ImgUrls = processUrls(c, act.ImgUrls)
	}
	return res
}
