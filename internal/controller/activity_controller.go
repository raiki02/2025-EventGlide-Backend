package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
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

	FindActBySearches() gin.HandlerFunc
	FindActByName() gin.HandlerFunc
	ShowActDetails() gin.HandlerFunc
	FindActByDate() gin.HandlerFunc
}

type ActController struct {
	ad   *dao.ActDao
	jwth *middleware.Jwt
	ch   *cache.Cache
	iu   *service.ImgUploader
	as   *service.ActivityService
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
// @Summary 通过bid查找活动
// @Produce json
// @Param bid formData string true "绑定id"
// @Success 200 {object} resp.Resp
// @Router /act/details [post]
func (ac ActController) ShowActDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		bid := c.PostForm("bid")
		if bid == "" {
			tools.ReturnMSG(c, "query cannot be nil", nil)
			return
		}
		as, err := ac.as.ShowActDetails(c, bid)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "success", as)
	}
}

// @Tags Activity
// @Summary 通过搜索条件查找活动
// @Produce json
// @Param actSearchReq body req.ActSearchReq true "搜索条件"
// @Success 200 {object} resp.Resp
// @Router /act/search [post]
func (ac ActController) FindActBySearches() gin.HandlerFunc {
	return func(c *gin.Context) {
		var actReq req.ActSearchReq
		err := c.ShouldBindJSON(&actReq)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		as, err := ac.as.FindActBySearches(c, &actReq)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "success", as)
	}
}

// @Tags Activity
// @Summary 通过日期查找活动
// @Produce json
// @Param date query string true "日期"
// @Success 200 {object} resp.Resp
// @Router /act/date [get]
func (ac ActController) FindActByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 02-01
		d := c.Query("date")
		if d == "" {
			tools.ReturnMSG(c, "query cannot be nil", nil)
			return
		}
		as, err := ac.ad.FindActByDate(c, d)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}
