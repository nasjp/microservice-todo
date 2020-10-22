package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/nasjp/microservice-todo/pb"
	"github.com/nasjp/microservice-todo/service/todo/internal/memory"
	"github.com/nasjp/microservice-todo/service/todo/internal/todo"
)

type Handler struct {
	memory *memory.Memory
}

func NewHandler(todoMemory *memory.Memory) *Handler {
	return &Handler{
		memory: todoMemory,
	}
}

var _ pb.TodoRepositoryServer = (*Handler)(nil)

func (s *Handler) Create(ctx context.Context, r *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	todo := &todo.TODO{
		ID:    uuid.New().String(),
		Title: r.GetTitle(),
	}

	if err := s.memory.Put(ctx, todo); err != nil {
		return nil, err
	}

	return &pb.CreateTodoResponse{Id: todo.ID}, nil
}

func (s *Handler) List(ctx context.Context, r *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	todos, err := s.memory.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Todo, 0, len(todos))

	for _, todo := range todos {
		res = append(res, &pb.Todo{Id: todo.ID, Title: todo.Title})
	}

	return &pb.ListTodoResponse{Todos: res}, nil
}
