package main

import (
	"fmt"
)

func climbStairs(n int) int {
	a := 1
	b := 1

	for i := 2; i < n+1; i++ {
		a,b = b,a+b
	}

	return b
}

func main() {
	result := climbStairs(1)
	fmt.Println(result)
	result = climbStairs(2)
	fmt.Println(result)
	result = climbStairs(3)
	fmt.Println(result)
	result = climbStairs(4)
	fmt.Println(result)
}
