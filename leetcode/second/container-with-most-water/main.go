// https://leetcode.cn/problems/container-with-most-water/
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

func maxArea(height []int) int {
	l, r, res := 0, len(height)-1, 0
	for l < r {
		// 双指针左右夹逼计算可拿到最大值
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l++
		} else {
			res = max(res, height[r]*(r-l))
			r--
		}
	}
	return res
}

func main() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(result)
	result = maxArea([]int{1, 1})
	fmt.Println(result)
}
