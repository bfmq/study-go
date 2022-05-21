package main

var (
	s2 *singleton = &singleton{}
)

func get2() *singleton {
	return s2
}
