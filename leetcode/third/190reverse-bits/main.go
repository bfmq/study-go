package main

import (
	"fmt"
)

func reverseBits(num uint32) (result uint32) {
	for i := 0; i < 32; i++ {
		x := num & 1
		num >>= 1
		result <<= 1
		if x == 1 {
			result++
		}
	}
	return
}

func main() {
	result := reverseBits(0)
	fmt.Println(result)
	result = reverseBits(0)
	fmt.Println(result)
}
