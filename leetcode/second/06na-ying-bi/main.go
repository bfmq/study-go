package main

import "fmt"

func minCount(coins []int) int {
	var res int
	for _, coin := range coins {
		// 每次就拿2
		if coin%2 == 0 {
			// 偶数直接算次数
			res += coin / 2
		} else {
			// 奇数多拿一次
			res += coin/2 + 1
		}
	}
	return res
}

func main() {
	result := minCount([]int{4, 2, 1})
	fmt.Println(result)
	result = minCount([]int{2, 3, 10})
	fmt.Println(result)
}
