package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/nasjp/microservice-todo/service/api/internal/repository"
)

var _ http.Handler = (*ListTODOsHandler)(nil)

type ListTODOsHandler struct {
	todoRepostiory *repository.TODORepository
}

func NewListTODOsHandler(todoRepository *repository.TODORepository) *ListTODOsHandler {
	return &ListTODOsHandler{
		todoRepostiory: todoRepository,
	}
}

func (h *ListTODOsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoRepostiory.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
