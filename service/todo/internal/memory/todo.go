package memory

import (
	"context"
	"sync"

	"github.com/nasjp/microservice-todo/service/todo/internal/todo"
)

type Memory struct {
	db   map[string]*todo.TODO
	lock sync.RWMutex
}

func NewMemory(db map[string]*todo.TODO) *Memory {
	return &Memory{db: db}
}

func (m *Memory) Put(ctx context.Context, t *todo.TODO) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()

	return nil
}

func (m *Memory) GetAll(ctx context.Context) (todo.TODOs, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := make(todo.TODOs, 0, len(m.db))
	for _, t := range m.db {
		result = append(result, t)
	}

	return result, nil
}
