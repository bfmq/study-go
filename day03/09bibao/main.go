package main

import "fmt"

func f1(f func()) {
	fmt.Println("f1")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x, y)
}

func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}
func main() {
	// f1(f2)
}
