package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/cache"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/tools"
)

// TODO: find函数写成过滤器模式？
type ActControllerHdl interface {
	NewAct() gin.HandlerFunc
	NewDraft() gin.HandlerFunc

	FindAllActs() gin.HandlerFunc
	FindActByHost() gin.HandlerFunc
	FindActByType() gin.HandlerFunc
	FindActByLocation() gin.HandlerFunc
	FindActByIfSignup() gin.HandlerFunc
	FindActByIsForeign() gin.HandlerFunc

	FindActByTime() gin.HandlerFunc
	FindActByName() gin.HandlerFunc
}

type ActController struct {
	ad   dao.ActDaoHdl
	jwth middleware.ClaimsHdl
	ch   cache.CacheHdl
}

func NewActController(ad dao.ActDaoHdl, jwth middleware.ClaimsHdl, ch cache.CacheHdl) ActControllerHdl {
	return &ActController{
		ad:   ad,
		ch:   ch,
		jwth: jwth,
	}
}

func (ac ActController) NewAct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func (ac ActController) NewDraft() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (ac ActController) FindAllActs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		as, err := ac.ad.FindAllActs(ctx)
		if err != nil {
			tools.ReturnMSG(ctx, err.Error(), nil)
			return
		}
		tools.ReturnMSG(ctx, "success", as)
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
