// https://leetcode.cn/problems/plus-one/
package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	// 从后向前处理
	for j := len(digits) - 1; j >= 0; j-- {
		if digits[j] != 9 {
			// 当前位不等于9则自增后返回结束
			digits[j]++
			return digits
		} else {
			// 设置为0进入下次循环自动处理下一位数
			digits[j] = 0
		}
	}

	// 处理首数字自增后升位情况
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	result := plusOne([]int{1, 2, 3})
	fmt.Println(result)
	result = plusOne([]int{4, 3, 2, 1})
	fmt.Println(result)
	result = plusOne([]int{0})
	fmt.Println(result)
	result = plusOne([]int{9, 9, 9})
	fmt.Println(result)
}
