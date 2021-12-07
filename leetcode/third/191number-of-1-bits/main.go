package main

import (
	"fmt"
)

func hammingWeight(num uint32) (result int) {
	for i := 0; i < 32; i++ {
		x := num & 1
		num >>= 1
		if x == 1 {
			result++
		}
	}
	return
}

func main() {
	result := hammingWeight(00000000000000000000000000001011)
	fmt.Println(result)
	result = hammingWeight(00000000000000000000000010000000)
	fmt.Println(result)
	// result = hammingWeight(11111111111111111111111111111101)
	// fmt.Println(result)
}
