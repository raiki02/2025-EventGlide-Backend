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
		if act.StudentID == "" || act.Title == "" || act.Introduce == "" {
			c.JSON(200, tools.ReturnMSG(c, "act param empty", nil))
			return
		}
		if len(act.LabelForm.Signer) <= 4 {
			c.JSON(200, tools.ReturnMSG(c, "signers at least 5", nil))
			return
		}

		if act.LabelForm.StartTime == "" || act.LabelForm.EndTime == "" {
			c.JSON(200, tools.ReturnMSG(c, "time error", nil))
			return
		}

		if act.LabelForm.EndTime < act.LabelForm.StartTime {
			c.JSON(200, tools.ReturnMSG(c, "start time greater than end time", nil))
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
// @Success 200 {object} resp.Resp{data=resp.CreateActivityResp}
// @Router /act/load [get]
func (ac *ActController) LoadDraft() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := tools.GetSid(c)
		if sid == "" {
			c.JSON(200, tools.ReturnMSG(c, "param empty", nil))
			return
		}
		d, err := ac.as.LoadDraft(c, sid)
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
// @Param name body req.FindActByNameReq true "活动名称"
// @Success 200 {object} resp.Resp{data=[]resp.ListActivitiesResp}
// @Router /act/name [post]
func (ac *ActController) FindActByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var r req.FindActByNameReq
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		if r.Name == "" {
			c.JSON(200, tools.ReturnMSG(c, "query cannot be nil", nil))
			return
		}
		as, err := ac.as.FindActByName(c, r.Name)
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
// @Param date body  req.FindActByDateReq true "日期查找"
// @Success 200 {object} resp.Resp{data=resp.ListActivitiesResp}
// @Router /act/date [post]
func (ac *ActController) FindActByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 02-01
		var r req.FindActByDateReq

		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}

		if r.Date == "" {
			c.JSON(200, tools.ReturnMSG(c, "query empty", nil))
			return
		}
		as, err := ac.as.FindActByDate(c, r.Date)
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
// @Param id path string true "用户id"
// @Success 200 {object} resp.Resp{data=resp.ListActivitiesResp}
// @Router /act/all/{id} [get]
func (ac *ActController) ListAllActs() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(200, tools.ReturnMSG(c, "query empty", nil))
			return
		}
		as, err := ac.as.ListAllActs(c, id)
		if err != nil {
			c.JSON(200, tools.ReturnMSG(c, err.Error(), nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", as))
	}
}
