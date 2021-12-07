package main

import "fmt"

func minCount(coins []int) int {
	result := 0
	for _, coin := range coins {
		if coin%2 == 0 {
			result += coin / 2
		} else {
			result += coin/2 + 1
		}
	}
	return result
}

func main() {
	result := minCount([]int{4, 2, 1})
	fmt.Println(result)
	result = minCount([]int{2, 3, 10})
	fmt.Println(result)
}
