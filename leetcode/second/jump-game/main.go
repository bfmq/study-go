// https://leetcode.cn/problems/jump-game/
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

func canJump(nums []int) bool {
	k := 0
	for i, num := range nums {
		if i > k {
			return false
		}
		k = max(k, i+num)
	}
	return true
}

func main() {
	result := canJump([]int{2, 3, 1, 1, 4})
	fmt.Println(result)
	result = canJump([]int{3, 2, 1, 0, 4})
	fmt.Println(result)

}
