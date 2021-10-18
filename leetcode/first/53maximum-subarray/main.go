package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) (result int) {
	if nums == nil {
		return
	}

	result = nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + int(math.Max(float64(nums[i-1]), 0))
		if nums[i] > result {
			result = nums[i]
		}
	}

	return
}

func main() {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(result)
	result = maxSubArray([]int{1})
	fmt.Println(result)
	result = maxSubArray([]int{0})
	fmt.Println(result)
	result = maxSubArray([]int{-1})
	fmt.Println(result)
	result = maxSubArray([]int{-100000})
	fmt.Println(result)
}
