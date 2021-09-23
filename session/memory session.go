package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	// 存kv
	data   map[string]interface{}
	rwLock sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	return &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
}

func (m *MemorySession) Set(key string, value interface{}) error {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.data[key] = value
	return nil
}

func (m *MemorySession) Get(key string) (interface{}, error) {
	m.rwLock.RLocker()
	defer m.rwLock.RUnlock()
	value, ok := m.data[key]
	if !ok {
		err := errors.New("key not exists in session")
		return nil, err
	}
	return value, nil
}

func (m *MemorySession) Del(key string) error {
	m.rwLock.RLocker()
	defer m.rwLock.RUnlock()
	delete(m.data, key)
	return nil
}

func (m *MemorySession) Save() error {
	return nil
}
