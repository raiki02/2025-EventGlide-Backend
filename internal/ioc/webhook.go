package ioc

import (
	"github.com/cqhasy/2025-Muxi-Team-auditor-Backend/sdk/webhook"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitListener(e *gin.Engine) *webhook.Listener {
	return &webhook.Listener{
		Engine: e,
		Path:   viper.GetString("auditor.webhookPath"),
	}
}
