package main

import "sync"

var (
	s3  *singleton
	mu3 sync.Mutex
)

func get3() *singleton {
	if s1 == nil {
		mu3.Lock()
		defer mu3.Unlock()
		if s1 == nil {
			s1 = &singleton{}
		}
	}
	return s3
}
