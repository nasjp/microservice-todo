package http

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	server *http.Server
}

type Config struct {
	Port                    int
	RepositoryTargetService string
	RepositoryTargetPort    int
}

func NewServer(cfg *Config) (*Server, error) {
	mux := http.NewServeMux()

	mux, err := routes(mux, cfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: mux,
		},
	}, nil
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to server: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}

	return nil
}
