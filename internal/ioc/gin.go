package ioc

import (
	"github.com/gin-gonic/gin"
)

func InitGinHandler() *gin.Engine {
	e := gin.New()
	e.Use(
		//gin.LoggerWithConfig(gin.LoggerConfig{
		//	SkipPaths: []string{"/metrics", "/"},
		//	Skip: func(c *gin.Context) bool {
		//		if strings.Contains(c.Request.RequestURI, "php") ||
		//			strings.Contains(c.Request.RequestURI, "favicon.ico") ||
		//			strings.Contains(c.Request.RequestURI, "css") ||
		//			strings.Contains(c.Request.RequestURI, "js") ||
		//			strings.Contains(c.Request.RequestURI, ".well-known") {
		//			return true
		//		}
		//		return false
		//	},
		//}),
		gin.Logger(),
		gin.Recovery(),
	)
	return e
}
