package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/nasjp/microservice-todo/service/api/internal/model"
	"github.com/nasjp/microservice-todo/service/api/internal/repository"
)

var _ http.Handler = (*CreateTODOHandler)(nil)

type CreateTODOHandler struct {
	todoRepostiory *repository.TODORepository
}

func NewCreateHandler(todoRepository *repository.TODORepository) *CreateTODOHandler {
	return &CreateTODOHandler{
		todoRepostiory: todoRepository,
	}
}

func (h *CreateTODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todo := &model.TODO{}

	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	newTODO, err := h.todoRepostiory.Create(todo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(newTODO); err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
