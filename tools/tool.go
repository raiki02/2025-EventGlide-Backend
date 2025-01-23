package tools

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/resp"
	"strconv"
)

func StrToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func Marshal(v interface{}) string {
	res, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(res)
}

func Unmarshal(data []byte, v interface{}) interface{} {
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil
	}
	return v
}

func ReturnMSG(c *gin.Context, msg string, res interface{}) resp.Resp {
	return resp.Resp{
		Code: c.Writer.Status(),
		Msg:  msg,
		Data: res,
	}
}
