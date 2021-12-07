package main

import (
	"fmt"
)

func removeDuplicates(nums []int) (result int) {
	if len(nums) == 0 {
		return
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	return len(nums)
}

func main() {
	result := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(result)
	result = removeDuplicates([]int{1, 1, 2})
	fmt.Println(result)
	result = removeDuplicates([]int{1, 1, 1, 1, 1, 1, 1, 1})
	fmt.Println(result)
}
