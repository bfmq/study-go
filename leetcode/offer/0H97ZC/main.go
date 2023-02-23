// https://leetcode.cn/problems/0H97ZC/
package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := make(map[int]int, len(arr2))
	for i, v := range arr2 {
		rank[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		if hasX && hasY {
			return rankX < rankY
		}
		if hasX || hasY {
			return hasX
		}
		return x < y
	})
	return arr1
}
