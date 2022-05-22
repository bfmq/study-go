package main

import "sync"

var (
	s1  *singleton
	mu1 sync.Mutex
)

func get1() *singleton {
	mu1.Lock()
	defer mu1.Unlock()

	if s1 == nil {
		s1 = &singleton{}
	}
	return s1
}
