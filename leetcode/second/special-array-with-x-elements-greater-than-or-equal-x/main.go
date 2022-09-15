// https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/
package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, n := 1, len(nums); i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}

func main() {
	result := specialArray([]int{3, 5})
	fmt.Println(result)
	result = specialArray([]int{0, 0})
	fmt.Println(result)
}
