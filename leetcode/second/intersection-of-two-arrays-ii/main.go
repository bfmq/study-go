// https://leetcode.cn/problems/intersection-of-two-arrays-ii/
package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	// 减少内存分配
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	// 计数nums1
	numsMap := make(map[int]int, len(nums1))
	for _, n := range nums1 {
		numsMap[n]++
	}

	// 对比nums2
	res := make([]int, 0, len(nums1))
	for _, n := range nums2 {
		if numsMap[n] > 0 {
			res = append(res, n)
			numsMap[n]--
		}
	}
	return res
}

func main() {
	resultSlice := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println(resultSlice)
	resultSlice = intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(resultSlice)
}
