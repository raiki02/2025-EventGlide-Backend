package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

type ActControllerHdl interface {
	NewAct() gin.HandlerFunc
	NewDraft() gin.HandlerFunc

	FindActBySearches() gin.HandlerFunc
	FindActByName() gin.HandlerFunc
	FindActByDate() gin.HandlerFunc
}

type ActController struct {
	as *service.ActivityService
	iu *service.ImgUploader
}

func NewActController(as *service.ActivityService, iu *service.ImgUploader) *ActController {
	return &ActController{
		as: as,
		iu: iu,
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
		//host,location,startTime,endTime,ifRegister,images,name
		err := c.ShouldBindJSON(&act)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}

		//处理用户上传图像给图床，返回储存url
		urls, err := ac.iu.ProcessImg(c)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		act.ImgUrls = urls

		err = ac.as.NewAct(c, &act)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}

		c.JSON(200, tools.ReturnMSG(c, "success", act))
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
		err = ac.as.NewDraft(c, d)
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
			c.JSON(200, tools.ReturnMSG(c, "query cannot be nil", nil))
			return
		}
		as, err := ac.as.FindActByName(c, n)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
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
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		as, err := ac.as.FindActBySearches(c, &actReq)
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
// @Param date query string true "日期"
// @Success 200 {object} resp.Resp
// @Router /act/date [get]
func (ac ActController) FindActByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 02-01
		d := c.Query("date")
		if d == "" {
			c.JSON(200, tools.ReturnMSG(c, "query empty", nil))
			return
		}
		as, err := ac.as.FindActByDate(c, d)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}
