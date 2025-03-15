package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/service"
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
