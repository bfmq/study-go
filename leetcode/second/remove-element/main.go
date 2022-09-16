// https://leetcode.cn/problems/remove-element/
package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	l := 0
	// 一遍循环一遍重新设置
	for _, v := range nums {
		if v != val {
			nums[l] = v
			l++
		}
	}
	return l
}

func main() {
	result := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(result)
	result = removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(result)
}
