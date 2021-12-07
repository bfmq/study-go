package main

import (
	"fmt"
)

func plusOne(digits []int) (result []int) {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] <= 9 {
			return digits
		} else {
			digits[i] = 0
		}
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
		return digits
	}
	return
}

func main() {
	result := plusOne([]int{1, 2, 3})
	fmt.Println(result)
	result = plusOne([]int{4, 3, 2, 1})
	fmt.Println(result)
	result = plusOne([]int{0})
	fmt.Println(result)
	result = plusOne([]int{9, 9, 9})
	fmt.Println(result)
}
