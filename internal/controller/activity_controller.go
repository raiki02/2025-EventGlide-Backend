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

// TODO: find函数写成过滤器模式
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

func (ac ActController) NewAct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var act model.Activity
		//获取用户填写信息
		err := ctx.ShouldBindJSON(&act)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}

		//处理用户上传图像给图床，返回储存url
		urls, err := ac.iu.ProcessImg(ctx)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		act.Images = urls

		//创建首帖关联id
		err = act.SetBid(ctx)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		//防止重复创建活动
		if ac.ad.CheckExist(ctx, act) {
			tools.ReturnMSG(ctx, "error exist", nil)
			return
		} else {
			err := ac.ad.CreateAct(ctx, act)
			if err != nil {
				tools.ReturnMSG(ctx, err.Error(), nil)
				return
			}
			tools.ReturnMSG(ctx, "success", act)
		}
	}
}

// 草稿逻辑：保存未完成上传的活动填写信息，绑定用户id（只能调用自己草稿），在上传后应该销毁，无法再次被调用到
// 将draft传给前端，前端获取字段信息，自动填入表中
// 需要用户单独调用（加载草稿）
func (ac ActController) NewDraft() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var d model.ActivityDraft
		//获取用户填写信息
		err := ctx.ShouldBindJSON(&d)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}

		//直接创建，不管有没有类似的
		//不保存上传图片，考虑图床空间
		//不设置绑定id，不一定会发布
		err = ac.ad.CreateDraft(ctx, d)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
	}
}

func (ac ActController) FindActByHost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		target := ctx.Query("host")
		if target == "" {
			tools.ReturnMSG(ctx, "query cannot be nil", nil)
			return
		}

		as, err := ac.ad.FindActByHost(ctx, target)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
	}
}

func (ac ActController) FindActByType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		target := ctx.Query("type")
		if target == "" {
			tools.ReturnMSG(ctx, "query cannot be nil", nil)
			return
		}

		as, err := ac.ad.FindActByType(ctx, target)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
	}
}

func (ac ActController) FindActByLocation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		target := ctx.Query("type")
		if target == "" {
			tools.ReturnMSG(ctx, "query cannot be nil", nil)
			return
		}

		as, err := ac.ad.FindActByLocation(ctx, target)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
	}
}

func (ac ActController) FindActByIfSignup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		target := ctx.Query("type")
		if target == "" {
			tools.ReturnMSG(ctx, "query cannot be nil", nil)
			return
		}

		as, err := ac.ad.FindActByIfSignup(ctx, target)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
	}
}

func (ac ActController) FindActByIsForeign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		target := ctx.Query("type")
		if target == "" {
			tools.ReturnMSG(ctx, "query cannot be nil", nil)
			return
		}

		as, err := ac.ad.FindActByIsForeign(ctx, target)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
	}
}

func (ac ActController) FindActByTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//format: yyyy-mm-dd hh:mm:ss in db
		start := ctx.Query("start_time") + ":00"
		end := ctx.Query("end_time") + ":00"

		as, err := ac.ad.FindActByTime(ctx, start, end)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
	}
}

func (ac ActController) FindActByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		n := ctx.Query("name")
		as, err := ac.ad.FindActByName(ctx, n)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
	}
}

func (ac ActController) FindActByDate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ../act/01-02 -> 01-02
		date := ctx.Param("date")
		as, err := ac.ad.FindActByDate(ctx, date)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
	}
}
