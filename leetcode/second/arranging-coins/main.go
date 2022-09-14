// https://leetcode.cn/problems/arranging-coins/
package main

import "fmt"

// 硬币够就循环计算下一层及剩余硬币
func arrangeCoins(n int) int {
	var res int
	for n > res {
		res++
		n -= res
	}
	return res
}

func main() {
	result := arrangeCoins(5)
	fmt.Println(result)
	result = arrangeCoins(8)
	fmt.Println(result)
}
