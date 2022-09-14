// https://leetcode.cn/problems/guess-number-higher-or-lower/
package main

import "fmt"

func guess(num int) int

// 普通二分查找对比即可
func guessNumber(n int) int {
	min, max := 0, n
	for min <= max {
		mid := min + (max-min)>>1
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == 1 {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return min
}

func main() {
	result := guessNumber(16)
	fmt.Println(result)
	result = guessNumber(14)
	fmt.Println(result)
	result = guessNumber(5)
	fmt.Println(result)
}
