package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/service"
)

type NumberControllerHdl interface {
	AddLikesNum() gin.HandlerFunc
	CutLikesNum() gin.HandlerFunc
	AddCommentsNum() gin.HandlerFunc
}

type NumberController struct {
	ns *service.NumberService
}

func NewNumberController(ns *service.NumberService) *NumberController {
	return &NumberController{
		ns: ns,
	}
}

func (nc *NumberController) AddLikesNum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var nr req.NumberReq
		if err := c.ShouldBindJSON(&nr); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = nc.ns.AddLikesNum(c, &nr)
	}
}

func (nc *NumberController) CutLikesNum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var nr req.NumberReq
		if err := c.ShouldBindJSON(&nr); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = nc.ns.CutLikesNum(c, &nr)
	}
}

func (nc *NumberController) AddCommentsNum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var nr req.NumberReq
		if err := c.ShouldBindJSON(&nr); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = nc.ns.AddCommentsNum(c, &nr)
	}
}
