// https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/
package main

import (
	"fmt"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	abs := func(a, b int) int {
		x := a - b
		if x < 0 {
			x *= -1
		}
		return x
	}
	var res int
	for i := 0; i < len(arr1); i++ {
		flag := true
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i], arr2[j]) <= d {
				flag = false
			}
		}
		if flag {
			res++
		}
	}
	return res
}

func main() {
	result := findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2)
	fmt.Println(result)
	result = findTheDistanceValue([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3)
	fmt.Println(result)
	result = findTheDistanceValue([]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6)
	fmt.Println(result)
}
