// https://leetcode.cn/problems/kth-missing-positive-number/
package main

import (
	"fmt"
)

func findKthPositive(arr []int, k int) int {
	// 原始列表与arr无关，永远是[1-1000]
	// arr[i]可以拆解成i+1+arr[i]前缺少数的数量
	// 缺少的第k个数就可以拆陈其前面arr[i]的i+1+k
	// 当arr[i]小于第k个，说明该数还没有没枚举到
	for _, v := range arr {
		if v <= k {
			k++
		} else {
			return k
		}
	}
	return k
}

func main() {
	result := findKthPositive([]int{2, 3, 4, 7, 11}, 5)
	fmt.Println(result)
	result = findKthPositive([]int{1, 2, 3, 4}, 2)
	fmt.Println(result)
}
