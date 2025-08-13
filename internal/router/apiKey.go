package router

import "github.com/cqhasy/2025-Muxi-Team-auditor-Backend/sdk/keyget"

type ApiKeyRouter interface {
	RegisterApiKeyRouters()
}
type apiKeyRouter struct {
	kg *keyget.KeyGet
}

func NewApiKeyRouter(kg *keyget.KeyGet) ApiKeyRouter {
	return &apiKeyRouter{kg: kg}
}

func (a *apiKeyRouter) RegisterApiKeyRouters() {
}
