// https://leetcode.cn/problems/climbing-stairs/
package main

import (
	"fmt"
)

func climbStairs(n int) int {
	// 左，右，合
	// 每次循环各自递增即可
	l, r, all := 0, 0, 1
	for i := 1; i <= n; i++ {
		l = r
		r = all
		all = l + r
	}
	return all
}

func main() {
	result := climbStairs(1)
	fmt.Println(result)
	result = climbStairs(2)
	fmt.Println(result)
	result = climbStairs(3)
	fmt.Println(result)
	result = climbStairs(4)
	fmt.Println(result)
}
