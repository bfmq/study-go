package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) (result []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}
	return
}

func main() {
	result := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println(result)
	result = intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(result)
}
