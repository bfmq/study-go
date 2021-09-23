package main

import "fmt"

func sum(x int, y int) (ret int) {
	return x + y
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)
}
