// https://leetcode.cn/problems/reverse-integer/
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var ans int
	for x != 0 {
		if ans > math.MaxInt32/10 || ans < math.MinInt32/10 {
			return 0
		}
		// 余数
		digit := x % 10
		// 返回值*10+余数
		ans = ans*10 + digit
		// 减位
		x /= 10
	}
	return ans
}

func main() {
	result := reverse(123)
	fmt.Println(result)
	result = reverse(-123)
	fmt.Println(result)
	result = reverse(120)
	fmt.Println(result)
	result = reverse(0)
	fmt.Println(result)
}
