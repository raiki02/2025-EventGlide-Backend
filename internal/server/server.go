package server

import (
	"github.com/google/wire"
	"github.com/raiki02/EG/internal/router"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	NewServer,
)

type Server struct {
	r *router.Router
	l *zap.Logger
}

func NewServer(r *router.Router, l *zap.Logger) *Server {
	return &Server{
		r: r,
		l: l,
	}
}

func (s *Server) Run() error {
	s.r.RegisterRouters()
	return s.r.Run()
}
