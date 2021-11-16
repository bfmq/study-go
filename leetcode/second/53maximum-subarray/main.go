package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) (result int) {
	// 入参检测
	result = nums[0]
	if len(nums) == 1 {
		return
	}

	for i := 1; i < len(nums); i++ {
		// 比较之前前一个数可达到的最大值与0比较大小，确定本值是否可增益
		nums[i] = nums[i] + max(0, nums[i-1])
		// 比较该下标最大值与原始最大值
		result = max(nums[i], result)
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
