package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/api/resp"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

type NumberControllerHdl interface {
	Send() gin.HandlerFunc
	Delete() gin.HandlerFunc
	Search() gin.HandlerFunc
}

type NumberController struct {
	ns *service.NumberService
}

func NewNumberController(ns *service.NumberService) *NumberController {
	return &NumberController{
		ns: ns,
	}
}

// @Tags Number
// @Summary Send a inteaction
// @Param req body req.NumberSendReq true "NumberSendReq"
// @Success 200 {object} resp.Resp
// @Router /number/create [post]
func (nc *NumberController) Send() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rq req.NumberSendReq
		err := c.ShouldBindJSON(&rq)
		if err != nil {
			c.JSON(400, tools.ReturnMSG(c, "bind error", nil))
			return
		}
		err = nc.ns.Send(c, rq)
		if err != nil {
			c.JSON(400, tools.ReturnMSG(c, "send error", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", nil))
	}
}

// @Tags Number
// @Summary Delete a inteaction
// @Param req body req.NumberDelReq true "NumberDelReq"
// @Success 200 {object} resp.Resp
// @Router /number/delete [post]
func (nc *NumberController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rq req.NumberDelReq
		err := c.ShouldBindJSON(&rq)
		if err != nil {
			c.JSON(400, tools.ReturnMSG(c, "bind error", nil))
			return
		}
		err = nc.ns.Delete(c, rq)
		if err != nil {
			c.JSON(400, tools.ReturnMSG(c, "del error", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", nil))
	}
}

// @Tags Number
// @Summary Search inteactions
// @Param req body req.NumberSearchReq true "NumberSearchReq"
// @Success 200 {object} resp.Resp{data=resp.NumberSearchResp}
// @Router /number/search [post]
func (nc *NumberController) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rq req.NumberSearchReq
		err := c.ShouldBindJSON(&rq)
		if err != nil {
			c.JSON(400, tools.ReturnMSG(c, "bind error", nil))
			return
		}
		numbers, count, err := nc.ns.Search(c, rq)
		if err != nil {
			c.JSON(400, tools.ReturnMSG(c, "search error", nil))
			return
		}
		c.JSON(200, tools.ReturnMSG(c, "success", resp.NumberSearchResp{
			Nums:  numbers,
			Total: count,
		}))
	}
}
