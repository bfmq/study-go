// https://leetcode.cn/problems/search-insert-position/
package main

import "fmt"

func searchInsert(nums []int, target int) int {
	// 取左右下标
	l, r := 0, len(nums)-1
	for l <= r {
		// 取中数
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r + 1
}

func main() {
	result := searchInsert([]int{1, 3, 5, 6}, 5)
	fmt.Println(result)
	result = searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Println(result)
	result = searchInsert([]int{1, 3, 5, 6}, 7)
	fmt.Println(result)
}
