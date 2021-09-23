package session

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]*Session
	rwLock     sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]*Session, 1024),
	}
}

func (m *MemorySessionMgr) Init(addr string, options ...string) error {
	return nil
}

func (m *MemorySessionMgr) Create() (Session, error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	u := uuid.NewV4()
	sessionId := u.String()
	s := NewMemorySession(sessionId)
	return s, nil
}

func (m *MemorySessionMgr) Get(sessionId string) (Session, error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	session, ok := m.sessionMap[sessionId]
	if !ok {
		err := errors.New("session not exists")
		return nil, err
	}

	return *session, nil
}
