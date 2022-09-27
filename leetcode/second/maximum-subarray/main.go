// https://leetcode.cn/problems/maximum-subarray/
package main

import "fmt"

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 原地修改改下标可获取的最大值
		if nums[i] < nums[i]+nums[i-1] {
			nums[i] = nums[i] + nums[i-1]
		}
		// 用修改后的值进行比较
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(result)
	result = maxSubArray([]int{1})
	fmt.Println(result)
	result = maxSubArray([]int{5, 4, -1, 7, 8})
	fmt.Println(result)
}
