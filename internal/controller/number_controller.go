package controller

import (
	"github.com/gin-gonic/gin"
	req "github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

type NumberControllerHdl interface {
	SendLikesNum() gin.HandlerFunc
	SendCommentsNum() gin.HandlerFunc
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
// @Summary 点赞数控制
// @Description not finished
// @Accept json
// @Produce json
// @Param likes_num body req.NumberReq true "点赞数"
// @Success 200 {object} resp.Resp
// @Router /number/likes [post]
func (nc *NumberController) SendLikesNum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var nq req.NumberReq
		if err := c.ShouldBindJSON(&nq); err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
		}
		nc.ns.SendLikesNum(c, &nq)
	}
}

// @Tags Number
// @Summary 评论数控制
// @Description not finished
// @Accept json
// @Produce json
// @Param comments_num body req.NumberReq true "评论数"
// @Success 200 {object} resp.Resp
// @Router /number/comments [post]
func (nc *NumberController) SendCommentsNum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var nq req.NumberReq
		if err := c.ShouldBindJSON(&nq); err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
		}
		nc.ns.SendCommentsNum(c, &nq)
	}
}
