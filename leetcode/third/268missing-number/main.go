package main

import "fmt"

func missingNumber(nums []int) (result int) {
	for i, num := range nums {
		result += (num - i)
	}
	result = len(nums) - result
	return
}

func main() {
	result := missingNumber([]int{3, 0, 1})
	fmt.Println(result)
	result = missingNumber([]int{0, 1})
	fmt.Println(result)
	result = missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	fmt.Println(result)
	result = missingNumber([]int{0})
	fmt.Println(result)
}
