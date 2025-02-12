package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/cache"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

// @TODO: find函数写成过滤器模式
type ActControllerHdl interface {
	NewAct() gin.HandlerFunc
	NewDraft() gin.HandlerFunc

	FindActByHost() gin.HandlerFunc
	FindActByType() gin.HandlerFunc
	FindActByLocation() gin.HandlerFunc
	FindActByIfSignup() gin.HandlerFunc
	FindActByIsForeign() gin.HandlerFunc

	FindActByTime() gin.HandlerFunc
	FindActByName() gin.HandlerFunc
	FindActByDate() gin.HandlerFunc
}

type ActController struct {
	ad   *dao.ActDao
	jwth *middleware.Jwt
	ch   *cache.Cache
	iu   *service.ImgUploader
}

func NewActController(ad *dao.ActDao, jwth *middleware.Jwt, ch *cache.Cache, iu *service.ImgUploader) *ActController {
	return &ActController{
		ad:   ad,
		ch:   ch,
		jwth: jwth,
		iu:   iu,
	}
}

// @Tags Activity
// @Summary 创建活动
// @Produce json
// @Accept json
// @Param activity body model.Activity true "活动"
// @Success 200 {object} resp.Resp
// @Router /act/create [post]
func (ac ActController) NewAct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var act model.Activity
		//获取用户填写信息
		err := c.ShouldBindJSON(&act)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}

		//处理用户上传图像给图床，返回储存url
		urls, err := ac.iu.ProcessImg(c)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		act.Images = urls

		//创建首帖关联id
		err = act.SetBid(c)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		//防止重复创建活动
		if ac.ad.CheckExist(c, act) {
			tools.ReturnMSG(c, "error exist", nil)
			return
		} else {
			err := ac.ad.CreateAct(c, act)
			if err != nil {
				tools.ReturnMSG(c, err.Error(), nil)
				return
			}
			tools.ReturnMSG(c, "success", act)
		}
	}
}

// @草稿逻辑：保存未完成上传的活动填写信息，绑定用户id（只能调用自己草稿），在上传后应该销毁，无法再次被调用到
// @将draft传给前端，前端获取字段信息，自动填入表中
// @需要用户单独调用（加载草稿）
// @Tags Activity
// @Summary 创建活动草稿
// @Description not finished
// @Produce json
// @Accept json
// @Param draft body model.ActivityDraft true "活动草稿"
// @Success 200 {object} resp.Resp
// @Router /act/draft [post]
func (ac ActController) NewDraft() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d model.ActivityDraft
		//获取用户填写信息
		err := c.ShouldBindJSON(&d)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}

		//直接创建，不管有没有类似的
		//不保存上传图片，考虑图床空间
		//不设置绑定id，不一定会发布
		err = ac.ad.CreateDraft(c, d)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
	}
}

// @Tags Activity
// @Summary 通过主办方查找活动
// @Produce json
// @Param host query string true "主办方"
// @Success 200 {object} resp.Resp
// @Router /act/host [get]
func (ac ActController) FindActByHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := c.Query("host")
		if target == "" {
			c.JSON(200, tools.ReturnMSG(c, "query cannot be nil", nil))
			return
		}

		as, err := ac.ad.FindActByHost(c, target)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}

// @Tags Activity
// @Summary 通过类型查找活动
// @Produce json
// @Param type query string true "类型"
// @Success 200 {object} resp.Resp
// @Router /act/type [get]
func (ac ActController) FindActByType() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := c.Query("type")
		if target == "" {
			c.JSON(200, tools.ReturnMSG(c, "query cannot be nil", nil))
			return
		}

		as, err := ac.ad.FindActByType(c, target)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}

// @Tags Activity
// @Summary 通过地点查找活动
// @Produce json
// @Param location query string true "地点"
// @Success 200 {object} resp.Resp
// @Router /act/location [get]
func (ac ActController) FindActByLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := c.Query("type")
		if target == "" {
			c.JSON(200, tools.ReturnMSG(c, "query cannot be nil", nil))
			return
		}

		as, err := ac.ad.FindActByLocation(c, target)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}

// @Tags Activity
// @Summary 通过是否需要报名查找活动
// @Produce json
// @Param type query string true "类型"
// @Success 200 {object} resp.Resp
// @Router /act/signup [get]
func (ac ActController) FindActByIfSignup() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := c.Query("type")
		if target == "" {
			c.JSON(200, tools.ReturnMSG(c, "query cannot be nil", nil))
			return
		}

		as, err := ac.ad.FindActByIfSignup(c, target)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}

// @Tags Activity
// @Summary 通过是否为外部活动查找活动
// @Produce json
// @Param type query string true "类型"
// @Success 200 {object} resp.Resp
// @Router /act/foreign [get]
func (ac ActController) FindActByIsForeign() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := c.Query("type")
		if target == "" {
			c.JSON(200, tools.ReturnMSG(c, "query cannot be nil", nil))
			return
		}

		as, err := ac.ad.FindActByIsForeign(c, target)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}

// @Tags Activity
// @Summary 通过时间查找活动
// @Produce json
// @Param start_time query string true "开始时间"
// @Param end_time query string true "结束时间"
// @Success 200 {object} resp.Resp
// @Router /act/time [get]
func (ac ActController) FindActByTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//format: yyyy-mm-dd hh:mm:ss in db
		start := c.Query("start_time") + ":00"
		end := c.Query("end_time") + ":00"
		if start == "" || end == "" {
			c.JSON(200, tools.ReturnMSG(c, "query cannot be nil", nil))
			return
		}

		as, err := ac.ad.FindActByTime(c, start, end)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}

// @Tags Activity
// @Summary 通过名称查找活动
// @Produce json
// @Param name query string true "名称"
// @Success 200 {object} resp.Resp
// @Router /act/name [get]
func (ac ActController) FindActByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Query("name")
		if n == "" {
			tools.ReturnMSG(c, "query cannot be nil", nil)
			return
		}
		as, err := ac.ad.FindActByName(c, n)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}

// @Tags Activity
// @Summary 通过日期查找活动
// @Produce json
// @Param date path string true "日期"
// @Success 200 {object} resp.Resp
// @Router /act/date/{date} [get]
func (ac ActController) FindActByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// @../act/01-02 -> 01-02
		date := c.Param("date")
		if date == "" {
			tools.ReturnMSG(c, "query cannot be nil", nil)
			return
		}
		as, err := ac.ad.FindActByDate(c, date)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}
