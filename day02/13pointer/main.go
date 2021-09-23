package main

import (
	"fmt"
)

func main() {
	// n := 18
	// fmt.Println(&n)
	// p := &n
	// fmt.Printf("%T\n", p)
	// m := *p
	// fmt.Println(m)

	var a = new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
}
