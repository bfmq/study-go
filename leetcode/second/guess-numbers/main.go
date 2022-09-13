// https://leetcode.cn/problems/guess-numbers/
package main

import "fmt"

func game(guess []int, answer []int) int {
	// guess跟answer是2人已经各自的答案顺序
	// 同下标为同一次游戏
	// 因此直接比较值是否相等即可知道本次游戏是否猜对
	// 最终将猜对累加即可
	var res int
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			res++
		}
	}
	return res
}

func main() {
	resultSlice := game([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(resultSlice)
}
