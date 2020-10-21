package http

import (
	"fmt"
	"net/http"

	"github.com/nasjp/microservice-todo/service/api/internal/handler"
	"github.com/nasjp/microservice-todo/service/api/internal/repository"
	"google.golang.org/grpc"
)

func routes(mux *http.ServeMux, cfg *Config) (*http.ServeMux, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.RepositoryTargetService, cfg.RepositoryTargetPort), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	mux.Handle("/create", handler.NewCreateHandler(repository.NewTODORepository(conn)))
	mux.Handle("/list", handler.NewListTODOsHandler(repository.NewTODORepository(conn)))

	return mux, nil
}
