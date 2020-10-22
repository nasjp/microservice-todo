package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	addr   string
	server *grpc.Server
}

func NewServer(port int) *Server {
	s := grpc.NewServer()
	register(s)

	return &Server{
		server: s,
		addr:   fmt.Sprintf(":%d", port),
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to server: %w", err)
	}

	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to server: %w", err)
	}

	return nil
}

func (s *Server) Stop() error {
	s.server.Stop()

	return nil
}
