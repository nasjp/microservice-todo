package grpc

import (
	"github.com/nasjp/microservice-todo/pb"
	"github.com/nasjp/microservice-todo/service/todo/internal/handler"
	"github.com/nasjp/microservice-todo/service/todo/internal/memory"
	"github.com/nasjp/microservice-todo/service/todo/internal/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func register(s *grpc.Server) {
	pb.RegisterTodoRepositoryServer(s, handler.NewHandler(memory.NewMemory(make(map[string]*todo.TODO))))

	reflection.Register(s)
}
