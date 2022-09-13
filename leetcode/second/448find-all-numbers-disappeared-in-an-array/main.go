package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) (result []int) {
	nList := make([]bool, len(nums)+1)
	nList[0] = true
	for _, num := range nums {
		nList[num] = true
	}
	for i, ok := range nList {
		if !ok {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	result := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(result)
	result = findDisappearedNumbers([]int{1, 1})
	fmt.Println(result)
}
