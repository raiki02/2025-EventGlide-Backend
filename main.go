//go:generate swag init && wire gen .
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/config"
	"github.com/spf13/viper"
	"log"
)

// @Title EventGlide API
// @Description 校灵通 API 文档
// @verstion 1.0
func main() {
	config.Init()
	fmt.Println(viper.GetString("kafka.addr"))
	e := gin.Default()
	app := InitApp(e)
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
