package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/nasjp/microservice-todo/pb"
	"github.com/nasjp/microservice-todo/service/repository/internal/memory"
)

type TODOHandler struct {
	todoMemory *memory.TODOMemory
}

func NewTODOHandler(todoMemory *memory.TODOMemory) *TODOHandler {
	return &TODOHandler{
		todoMemory: todoMemory,
	}
}

var _ pb.TodoRepositoryServer = (*TODOHandler)(nil)

func (s *TODOHandler) Create(ctx context.Context, r *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	todo := &memory.TODO{
		ID:    uuid.New().String(),
		Title: r.GetTitle(),
	}

	if err := s.todoMemory.PutTODO(ctx, todo); err != nil {
		return nil, err
	}

	return &pb.CreateTodoResponse{Id: todo.ID}, nil
}

func (s *TODOHandler) List(ctx context.Context, r *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	todos, err := s.todoMemory.GetAllTODOs(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Todo, 0, len(todos))

	for _, todo := range todos {
		res = append(res, &pb.Todo{Id: todo.ID, Title: todo.Title})
	}

	return &pb.ListTodoResponse{Todos: res}, nil
}
