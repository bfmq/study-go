package main

import (
	"fmt"
)

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + i&1
	}
	return result
}

func main() {
	result := countBits(2)
	fmt.Println(result)
	result = countBits(5)
	fmt.Println(result)
}
