// https://leetcode.cn/problems/binary-search/
package main

import "fmt"

// 普通二分查找对比即可
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

func main() {
	result := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(result)
	result = search([]int{-1, 0, 3, 5, 9, 12}, 2)
	fmt.Println(result)
	result = search([]int{5}, 5)
	fmt.Println(result)
}
