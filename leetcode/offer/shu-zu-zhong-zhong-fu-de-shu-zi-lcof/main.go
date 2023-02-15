// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	nMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := nMap[v]; ok {
			return v
		} else {
			nMap[v] = true
		}
	}
	return -1
}

func main() {
	result := findRepeatNumber([]int{1, 2, 3, 4, 5, 5})
	fmt.Println(result)
}
