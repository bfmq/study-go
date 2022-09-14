// https://leetcode.cn/problems/valid-perfect-square/
package main

import "fmt"

func isPerfectSquare(num int) bool {
	min, max := 0, num
	for min <= max {
		mid := min + (max-min)>>1
		square := mid * mid
		if square < num {
			min = mid + 1
		} else if square > num {
			max = mid - 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	result := isPerfectSquare(16)
	fmt.Println(result)
	result = isPerfectSquare(14)
	fmt.Println(result)
	result = isPerfectSquare(5)
	fmt.Println(result)
}
