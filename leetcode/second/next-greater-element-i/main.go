// https://leetcode.cn/problems/next-greater-element-i/
package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}

	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = mp[num]
	}
	return res
}

func main() {
	result := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	fmt.Println(result)
	result = nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})
	fmt.Println(result)

}
