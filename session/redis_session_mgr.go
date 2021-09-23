package session

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

type RedisSessionMgr struct {
	addr       string
	password   string
	pool       *redis.Pool
	sessionMap map[string]*Session
	rwLock     sync.RWMutex
}

func NewRedisSessionMgr() *RedisSessionMgr {
	return &RedisSessionMgr{
		sessionMap: make(map[string]*Session, 1024),
	}
}

func myPool(addr, password string) (pool *redis.Pool) {
	return &redis.Pool{
		MaxIdle:   64,
		MaxActive: 1000,
		// IdleTimeout: time.Second,
		Dial: func() (*redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				fmt.Printf(",err:%s\n", err)
				return nil, err
			}
			_, err = conn.Do("AUTH", password)
			if err != nil {
				conn.Close()
				fmt.Printf(",err:%s\n", err)
				return nil, err
			}
			return conn, err
		},
	}
}

func (r *RedisSessionMgr) Init(addr string, options ...string) error {
	if len(options) > 0 {
		r.password = options[0]
	}

	return nil
}

func (r *RedisSessionMgr) Create() (Session, error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	u := uuid.NewV4()
	sessionId := u.String()
	s := NewRedisSession(sessionId, r.pool)
	return s, nil
}

func (r *RedisSessionMgr) Get(sessionId string) (Session, error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	session, ok := r.sessionMap[sessionId]
	if !ok {
		err := errors.New("session not exists")
		return nil, err
	}

	return *session, nil
}
