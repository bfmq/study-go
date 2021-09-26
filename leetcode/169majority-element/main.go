package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) (result int) {
	numsLength := len(nums)
	sort.Ints(nums)
	return nums[numsLength/2]
}

func main() {
	result := majorityElement([]int{3, 2, 3})
	fmt.Println(result)
	result = majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(result)
	result = majorityElement([]int{10})
	fmt.Println(result)
}
