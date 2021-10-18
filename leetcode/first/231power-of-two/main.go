package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo(n int) bool {
	if n < int(math.Pow(-2, 31)) || n > int(math.Pow(2, 31))-1 || n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}

	return n == 1
}

func main() {
	result := isPowerOfTwo(1)
	fmt.Println(result)
	result = isPowerOfTwo(16)
	fmt.Println(result)
	result = isPowerOfTwo(3)
	fmt.Println(result)
	result = isPowerOfTwo(4)
	fmt.Println(result)
	result = isPowerOfTwo(5)
	fmt.Println(result)
}
