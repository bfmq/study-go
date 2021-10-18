package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		} else {
			m[num] = true
		}
	}
	return false
}

func main() {
	result := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Println(result)
	result = containsDuplicate([]int{1, 2, 3, 4})
	fmt.Println(result)
	result = containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	fmt.Println(result)
}
