package rest

import (
	"context"
	"net/http"
	"strconv"
	"userLoggingProject/internal/core/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.LaunchConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + strconv.Itoa(cfg.Port),
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
