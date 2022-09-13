// https://leetcode.cn/problems/na-ying-bi/
package main

import "fmt"

func minCount(coins []int) int {
	var res int
	// 每次就拿2
	for _, coin := range coins {
		res += (coin + 1) / 2
	}
	return res
}

func main() {
	result := minCount([]int{4, 2, 1})
	fmt.Println(result)
	result = minCount([]int{2, 3, 10})
	fmt.Println(result)
}
