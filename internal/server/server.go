package server

import (
	"github.com/raiki02/EG/internal/router"
)

type Server struct {
	r *router.Router
}

func NewServer(r *router.Router) *Server {
	return &Server{r}
}

func (s *Server) Run() error {
	s.r.RegisterRouters()
	return s.r.Run()
}
