package main

import (
	"fmt"
)

func step(n int) int {
	var sum int
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	nMap := make(map[int]bool, 32)
	for ; n != 1 && !nMap[n]; n, nMap[n] = step(n), true {
	}
	return n == 1
}

func main() {
	result := isHappy(19)
	fmt.Println(result)
	result = isHappy(2)
	fmt.Println(result)
	result = isHappy(1)
	fmt.Println(result)
	result = isHappy(1111111)
	fmt.Println(result)
}
