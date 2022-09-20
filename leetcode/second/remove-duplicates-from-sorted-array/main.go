// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
package main

import (
	"fmt"
)

func removeDuplicates(nums []int) (result int) {
	// 从后向前处理不会影响元数组下标
	for j := len(nums) - 1; j > 0; j-- {
		if nums[j] == nums[j-1] {
			// 如果当前下标与前一个下标数字相同，则要被清理
			nums = append(nums[:j], nums[j+1:]...)
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
