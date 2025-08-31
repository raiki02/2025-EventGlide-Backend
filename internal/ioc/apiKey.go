package ioc

import (
	"github.com/cqhasy/2025-Muxi-Team-auditor-Backend/sdk/keyget"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitApiKeyGetter(e *gin.Engine) *keyget.KeyGet {
	return keyget.DefaultServe(e, "", viper.GetString("auditor.apiKeyPath"))
}
