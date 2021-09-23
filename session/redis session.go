package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

type RedisSession struct {
	sessionId string
	// 存kv
	sessionMap map[string]interface{}
	rwLock     sync.RWMutex
	pool       *redis.Pool
	flag       int
}

const (
	SessionFlagNone = iota
	SessionFlagModify
)

func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	return &RedisSession{
		sessionId:  id,
		sessionMap: make(map[string]interface{}, 16),
		pool:       pool,
		flag:       SessionFlagNone,
	}
}

func (r *RedisSession) Set(key string, value interface{}) error {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	r.sessionMap[key] = value
	r.flag = SessionFlagModify
	return nil
}

func (r *RedisSession) Get(key string) (interface{}, error) {
	r.rwLock.RLocker()
	defer r.rwLock.RUnlock()

	value, ok := r.sessionMap[key]
	if !ok {
		err := errors.New("key not exists in session")
		return nil, err
	}
	return value, nil
}

func (r *RedisSession) loadFromRedis(key string) error {
	coon := r.pool.Get()
	reply, err := coon.Do("GET", r.sessionId)
	if err != nil {
		fmt.Printf("save to redis failed,err:%s\n", err)
		return err
	}
	data, err := redis.String(reply, err)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisSession) Del(key string) error {
	r.rwLock.RLocker()
	defer r.rwLock.RUnlock()
	r.flag = SessionFlagModify

	delete(r.sessionMap, key)
	return nil
}

func (r *RedisSession) Save() error {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	if r.flag != SessionFlagModify {
		err := errors.New("don't need sava")
		return err
	}

	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		fmt.Printf(",err:%s\n", err)
		return err
	}
	coon := r.pool.Get()
	_, err = coon.Do("SET", r.sessionId, string(data))
	if err != nil {
		fmt.Printf("save to redis failed,err:%s\n", err)
		return err
	}
	r.flag = SessionFlagNone
	return nil
}
