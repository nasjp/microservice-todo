package memory

import (
	"context"
	"sync"
)

type TODO struct {
	ID    string
	Title string
}

type TODOs []*TODO

// var _ DB = (*memoryDB)(nil)

type TODOMemory struct {
	db   map[string]*TODO
	lock sync.RWMutex
}

func NewTODOMemory(db map[string]*TODO) *TODOMemory {
	return &TODOMemory{db: db}
}

func (m *TODOMemory) PutTODO(ctx context.Context, t *TODO) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()

	return nil
}

func (m *TODOMemory) GetAllTODOs(ctx context.Context) (TODOs, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := make(TODOs, 0, len(m.db))
	for _, t := range m.db {
		result = append(result, t)
	}

	return result, nil
}
