package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsHdl interface {
	HandleCors() error
}

type Cors struct {
	e *gin.Engine
}

func NewCors(e *gin.Engine) *Cors {
	return &Cors{e: e}
}

func (c *Cors) HandleCors() error {
	corsConf := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	c.e.Use(cors.New(corsConf))
	return nil
}
