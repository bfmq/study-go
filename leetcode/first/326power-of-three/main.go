package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(x int) (result bool) {
	if x < int(math.Pow(-2, 31)) || x > int(math.Pow(2, 31))-1 || x <= 0 {
		return false
	}

	for x%3 == 0 {
		x /= 3
	}

	return x == 1
}

func main() {
	result := isPowerOfThree(27)
	fmt.Println(result)
	result = isPowerOfThree(0)
	fmt.Println(result)
	result = isPowerOfThree(9)
	fmt.Println(result)
	result = isPowerOfThree(45)
	fmt.Println(result)
}
