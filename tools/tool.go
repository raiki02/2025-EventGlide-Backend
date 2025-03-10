package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func GenUUID() string {
	u, _ := uuid.NewUUID()
	return u.String()
}

func SliceToString(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.Join(s, ",")
}

func StringToSlice(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

func ReturnMSG(c *gin.Context, msg string, res interface{}) map[string]interface{} {
	return gin.H{
		"code": c.Writer.Status(),
		"msg":  msg,
		"data": res,
	}
}
