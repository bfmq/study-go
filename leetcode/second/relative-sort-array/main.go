// https://leetcode.cn/problems/relative-sort-array/
package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// 记录arr2值与下标对应关系
	rank := map[int]int{}
	for i, v := range arr2 {
		rank[v] = i
	}
	// 自定义排序比较函数
	sort.Slice(arr1, func(i, j int) bool {
		// 取出比较的2值
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		// 如果2值都存在于arr2，则对比下标
		if hasX && hasY {
			return rankX < rankY
		}
		// 只有一个情况返回hasX即可比较
		if hasX || hasY {
			return hasX
		}
		// 都不存在正常比较值大小
		return x < y
	})
	return arr1
}

func main() {
	result := relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	fmt.Println(result)
	result = relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6})
	fmt.Println(result)
}
