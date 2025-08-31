package tools

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"time"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

func GenUUID() string {
	var b strings.Builder
	for i := 0; i < 4; i++ {
		b.WriteRune(letterRunes[rand.Intn(len(letterRunes))])
	}
	return b.String()
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

func GetSid(c *gin.Context) string {
	sid, ok := c.Get("studentid")
	if !ok {
		return ""
	}
	res, ok := sid.(string)
	if !ok {
		return ""
	}
	return res
}

func ParseTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StatusMapper(auditStatus string) string {
	//  0: "未审核",
	//    1: "通过",
	//    2: "不通过",
	switch auditStatus {
	case "未审核":
		return "pending"
	case "通过":
		return "pass"
	case "不通过":
		return "reject"
	default:
		return "unknown error"
	}
}

func IfRegisterMapper(_if string) bool {
	return _if == "是"
}
