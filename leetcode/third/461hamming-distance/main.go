package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	result := hammingDistance(1, 4)
	fmt.Println(result)
	result = hammingDistance(3, 1)
	fmt.Println(result)
}
