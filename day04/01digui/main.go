package main

import "fmt"

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}

	return n * f(n-1)
}

func main() {
	// r1 := f(5)
	// fmt.Println(r1)
	r2 := taijie(5)
	fmt.Println(r2)
}
