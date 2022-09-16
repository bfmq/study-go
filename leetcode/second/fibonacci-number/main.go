// https://leetcode.cn/problems/fibonacci-number/
package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	// 从2开始迭代即可
	l, r, all := 0, 0, 1
	for i := 2; i <= n; i++ {
		l = r
		r = all
		all = l + r
	}
	return all
}

func main() {
	result := fib(2)
	fmt.Println(result)
	result = fib(3)
	fmt.Println(result)
	result = fib(4)
	fmt.Println(result)
}
