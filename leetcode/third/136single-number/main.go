package main

import (
	"fmt"
)

func singleNumber(nums []int) (x int) {
	for _, num := range nums {
		x ^= num
	}
	return
}

func main() {
	result := singleNumber([]int{2, 2, 1})
	fmt.Println(result)
	result = singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Println(result)
}
