package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/config"
	"log"
)

// @Title EventGlide API
// @Description 校灵通 API 文档
// @verstion 1.0
func main() {
	config.Init()
	e := gin.Default()
	app := InitApp(e)
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
