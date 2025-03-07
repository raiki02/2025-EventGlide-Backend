package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

type ActControllerHdl interface {
	NewAct() gin.HandlerFunc
	NewDraft() gin.HandlerFunc
	LoadDraft() gin.HandlerFunc
	FindActBySearches() gin.HandlerFunc
	FindActByName() gin.HandlerFunc
	FindActByDate() gin.HandlerFunc
	FindActByOwnerID() gin.HandlerFunc
	ListAllActs() gin.HandlerFunc
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
// @Param activity body req.CreateActReq true "活动"
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp{data=resp.CreateActivityResp}
// @Router /act/create [post]
func (ac *ActController) NewAct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var act req.CreateActReq
		//获取用户填写信息
		err := c.ShouldBindJSON(&act)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}

		a, err := ac.as.NewAct(c, &act)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}

		c.JSON(200, tools.ReturnMSG(c, "success", a))
	}
}

// @Tags Activity
// @Summary 创建活动草稿
// @Description not finished
// @Produce json
// @Accept json
// @Param draft body req.CreateActReq true "活动草稿"
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp{data=resp.CreateActivityResp}
// @Router /act/draft [post]
func (ac *ActController) NewDraft() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d req.CreateActReq
		//获取用户填写信息
		err := c.ShouldBindJSON(&d)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}

		res, err := ac.as.NewDraft(c, &d)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", res))
	}
}

// @Tags Activity
// @Summary 加载活动草稿
// @Produce json
// @Accept json
// @Param Authorization header string true "token"
// @Param draft body req.DraftReq true "加载草稿"
// @Success 200 {object} resp.Resp{data=resp.CreateActivityResp}
// @Router /act/load [post]
func (ac ActController) LoadDraft() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dReq req.DraftReq
		err := c.ShouldBindJSON(&dReq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		d, err := ac.as.LoadDraft(c, dReq)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", d))
	}
}

// @Tags Activity
// @Summary 通过名称查找活动
// @Produce json
// @Param Authorization header string true "token"
// @Param name path string true "名称"
// @Success 200 {object} resp.Resp{data=resp.ListActivitiesResp}
// @Router /act/name/{name} [get]
func (ac *ActController) FindActByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("name")
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
// @Param Authorization header string true "token"
// @Param actSearchReq body req.ActSearchReq true "搜索条件"
// @Success 200 {object} resp.Resp{data=resp.ListActivitiesResp}
// @Router /act/search [post]
func (ac *ActController) FindActBySearches() gin.HandlerFunc {
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
// @Param Authorization header string true "token"
// @Param date path string true "日期"
// @Success 200 {object} resp.Resp{data=resp.ListActivitiesResp}
// @Router /act/date/{date} [get]
func (ac *ActController) FindActByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 02-01
		d := c.Param("date")
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

// @Tags Activity
// @Summary 通过创建者id查找活动
// @Produce json
// @Param Authorization header string true "token"
// @Param id path string true "创建者id"
// @Success 200 {object} resp.Resp{data=resp.ListActivitiesResp}
// @Router /act/own/{id} [get]
func (ac *ActController) FindActByOwnerID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(200, tools.ReturnMSG(c, "query empty", nil))
			return
		}
		as, err := ac.as.FindActByOwnerID(c, id)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}

// @Tags Activity
// @Summary 列出所有活动
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp{data=resp.ListActivitiesResp}
// @Router /act/all [get]
func (ac *ActController) ListAllActs() gin.HandlerFunc {
	return func(c *gin.Context) {
		as, err := ac.as.ListAllActs(c)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}
