package repository

import (
	"context"

	"github.com/nasjp/microservice-todo/pb"
	"github.com/nasjp/microservice-todo/service/api/internal/model"
	"google.golang.org/grpc"
)

type TODORepository struct {
	client pb.TodoRepositoryClient
}

func NewTODORepository(conn *grpc.ClientConn) *TODORepository {
	return &TODORepository{
		client: pb.NewTodoRepositoryClient(conn),
	}
}

func (r *TODORepository) Create(todo *model.TODO) (*model.TODO, error) {
	res, err := r.client.Create(context.Background(), &pb.CreateTodoRequest{Title: todo.Title})
	if err != nil {
		return nil, err
	}

	return &model.TODO{
		ID:    res.GetId(),
		Title: todo.Title,
	}, nil
}

func (r *TODORepository) List() (model.TODOs, error) {
	res, err := r.client.List(context.Background(), &pb.ListTodoRequest{})
	if err != nil {
		return nil, err
	}

	todos := make(model.TODOs, 0, len(res.GetTodos()))

	for _, r := range res.GetTodos() {
		todos = append(todos, &model.TODO{ID: r.Id, Title: r.Title})
	}

	return todos, nil
}
