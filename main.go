package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/config"
	"log"
)

func main() {
	config.Init()
	e := gin.Default()
	app := InitApp(e)
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
