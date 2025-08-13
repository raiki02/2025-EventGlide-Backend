package service

import (
	"github.com/cqhasy/2025-Muxi-Team-auditor-Backend/api/request"
	"github.com/cqhasy/2025-Muxi-Team-auditor-Backend/sdk/webhook"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/tools"
	"go.uber.org/zap"
)

type CallbackAuditorService interface {
}

type callbackAuditorService struct {
	repo dao.AuditorRepository

	listener *webhook.Listener
	l        *zap.Logger
}

func NewCallbackAuditor(repo dao.AuditorRepository, listener *webhook.Listener, l *zap.Logger) CallbackAuditorService {
	// TODO: 这里太强耦合了, 但是这个handler必须要用repo没办法
	f := webhook.HandlerFunc(func(e string, p request.HookPayload) {
		if m, ok := p.Data.(map[string]interface{}); ok {
			id := uint(m["id"].(float64)) // JSON 数字会被解析成 float64
			status := tools.StatusMapper(m["status"].(string))
			ctx := gin.Context{}

			if err := repo.Update(&ctx, id, status); err != nil {
				l.Error("Failed to update auditor status", zap.Uint("id", id), zap.String("status", status), zap.Error(err))
			} else {
				l.Info("Auditor status updated successfully", zap.Uint("id", id), zap.String("status", status))
			}
		}
	})
	listener.Handler = f
	listener.RegisterRoutes()

	return &callbackAuditorService{
		repo:     repo,
		listener: listener,
		l:        l.Named("auditor/callback"),
	}
}
