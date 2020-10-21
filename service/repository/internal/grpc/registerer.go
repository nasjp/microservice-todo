package grpc

import (
	"github.com/nasjp/microservice-todo/pb"
	"github.com/nasjp/microservice-todo/service/repository/internal/handler"
	"github.com/nasjp/microservice-todo/service/repository/internal/memory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func register(s *grpc.Server) {
	pb.RegisterTodoRepositoryServer(s, handler.NewTODOHandler(memory.NewTODOMemory(make(map[string]*memory.TODO))))

	reflection.Register(s)
}
