package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/config"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/ioc"
	"github.com/raiki02/EG/internal/router"
	"time"
)

func main() {
	config.Init()
	r := gin.Default()
	db := ioc.InitDB()
	csvc := controller.NewCCNUService(time.Second * 5)
	dao := dao.NewUserDAO(db)
	uc := controller.NewUserController(r, dao, *csvc)
	ur := router.NewUserRouter(r, uc)
	ur.RegisterUserRouters()
	r.Run()
}
