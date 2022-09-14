// https://leetcode.cn/problems/sqrtx/
package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	min, max := 0, x
	// 可以执行平方根情况
	for max-min > 1 {
		// 中间数
		mid := (max + min) >> 1
		// 等价于mid*mid > x
		if x/mid < mid {
			max = mid
		} else {
			min = mid
		}
	}
	return min
}

func main() {
	result := mySqrt(4)
	fmt.Println(result)
	result = mySqrt(8)
	fmt.Println(result)
}
