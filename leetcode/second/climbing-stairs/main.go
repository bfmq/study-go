// https://leetcode.cn/problems/climbing-stairs/
package main

import (
	"fmt"
)

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
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
