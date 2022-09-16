// https://leetcode.cn/problems/n-th-tribonacci-number/
package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	
	a, b, c, all := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		a = b
		b = c
		c = all
		all = a + b + c
	}
	return all
}

func main() {
	result := tribonacci(4)
	fmt.Println(result)
	result = tribonacci(25)
	fmt.Println(result)
}
