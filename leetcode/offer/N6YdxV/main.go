// https://leetcode.cn/problems/N6YdxV/
package main

func searchInsert(nums []int, target int) (res int) {
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
	return l
}
