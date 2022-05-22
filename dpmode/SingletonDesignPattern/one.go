package main

import "sync"

var (
	s4   *singleton
	once sync.Once
)

func get4() *singleton {
	once.Do(func() {
		s4 = &singleton{}
	})
	return s4
}
