package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/cache"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
)

type ActivityControllerHdl interface {
	NewActivity() gin.HandlerFunc
	NewDraft() gin.HandlerFunc
	ListAllActivity() gin.HandlerFunc
	ListActivityByType() gin.HandlerFunc
	ListActivityByTime() gin.HandlerFunc
	ListActivityByHost() gin.HandlerFunc
	ListActivityByLocation() gin.HandlerFunc
	ListActivityByName() gin.HandlerFunc
}

type ActivityController struct {
	c   cache.CacheHdl
	adh dao.ActivityDaoHdl
}

func NewActivityController(c cache.CacheHdl, adh dao.ActivityDaoHdl) ActivityControllerHdl {
	return &ActivityController{c: c, adh: adh}
}

func (a *ActivityController) NewActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var activity model.Activity
		err := c.ShouldBindJSON(&activity)
		if err != nil {
			tools.ReturnMSG(c, "error", err)
			return
		}
		err = a.adh.NewActivity(c, activity)
		if err != nil {
			tools.ReturnMSG(c, "error", err)
			return
		}
		tools.ReturnMSG(c, "activity created", nil)
	}
}

func (a *ActivityController) NewDraft() gin.HandlerFunc {
	return func(c *gin.Context) {
		var activityDraft model.ActivityDraft
		err := c.ShouldBindJSON(&activityDraft)
		if err != nil {
			tools.ReturnMSG(c, "error", err)
			return
		}
		err = a.adh.NewDraft(c, activityDraft)
		if err != nil {
			tools.ReturnMSG(c, "error", err)
			return
		}
		tools.ReturnMSG(c, "draft created", nil)
	}
}

// todo 排序方式 持续几天的(结束更早的排在上面)和当天进行的(开始更早的排在上面)
func (a *ActivityController) ListAllActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		activities, err := a.adh.ListAllActivity(c)
		if err != nil {
			tools.ReturnMSG(c, "error", err)
			return
		}
		tools.ReturnMSG(c, "success", activities)
	}
}

func (a *ActivityController) ListActivityByType() gin.HandlerFunc {
	return func(c *gin.Context) {
		tp := c.Query("type")
		activities, err := a.c.Get(c, tp)
		if err == nil {
			tools.ReturnMSG(c, "get cache success", activities)
			return
		}
		activities, err = a.adh.ListActivityByType(c, tp)
		if err != nil {
			tools.ReturnMSG(c, "get dao error", err)
			return
		}
		err = a.c.Set(c, tp, activities)
		if err != nil {
			tools.ReturnMSG(c, "set cache error", err)
			return
		}
		tools.ReturnMSG(c, "get dao success", activities)
	}
}

// todo 时间这样找不了区间
func (a *ActivityController) ListActivityByTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		time := c.Query("time")
		//先去缓存
		activities, err := a.c.Get(c, time)
		if err == nil {
			tools.ReturnMSG(c, "success", activities)
			return
		}
		//不能因为缓存as为空而去数据库查找，因为可能数据库中没有这个数据

		//去数据库
		activities, err = a.adh.ListActivityByTime(c, time)
		if err != nil {
			tools.ReturnMSG(c, "error", err)
			return
		}
		err = a.c.Set(c, time, activities)
		if err != nil {
			tools.ReturnMSG(c, "error", err)
			return
		}
		tools.ReturnMSG(c, "success", activities)
	}
}

func (a *ActivityController) ListActivityByHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Query("host")
		activities, err := a.c.Get(c, host)
		if err == nil {
			tools.ReturnMSG(c, "get cache success", activities)
			return
		}
		activities, err = a.adh.ListActivityByHost(c, host)
		if err != nil {
			tools.ReturnMSG(c, "get dao error", err)
			return
		}
		err = a.c.Set(c, host, activities)
		if err != nil {
			tools.ReturnMSG(c, "set cache error", err)
			return
		}
		tools.ReturnMSG(c, "get dao success", activities)
	}
}

func (a *ActivityController) ListActivityByLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		location := c.Query("location")
		activities, err := a.c.Get(c, location)
		if err == nil {
			tools.ReturnMSG(c, "get cache success", activities)
			return
		}
		activities, err = a.adh.ListActivityByLocation(c, location)
		if err != nil {
			tools.ReturnMSG(c, "get dao error", err)
			return
		}
		err = a.c.Set(c, location, activities)
		if err != nil {
			tools.ReturnMSG(c, "set cache error", err)
			return
		}
		tools.ReturnMSG(c, "get dao success", activities)
	}
}

func (a *ActivityController) ListActivityByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		activities, err := a.c.Get(c, name)
		if err == nil {
			tools.ReturnMSG(c, "get cache success", activities)
			return
		}
		activities, err = a.adh.ListActivityByName(c, name)
		if err != nil {
			tools.ReturnMSG(c, "get dao error", err)
			return
		}
		err = a.c.Set(c, name, activities)
		if err != nil {
			tools.ReturnMSG(c, "set cache error", err)
			return
		}
		tools.ReturnMSG(c, "get dao success", activities)
	}
}
