package main

import (
	"fmt"
)

func keep(x int) bool {
	if x > 5 {
		return true
	}
	return false
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := 0
	for _, x := range a {
		if keep(x) {
			a[n] = x
			n++
		}
	}
	fmt.Println(a)
	a = a[:n]
	fmt.Println(a)
}
