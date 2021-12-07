package main

import "fmt"

func game(guess []int, answer []int) int {
	result := 0
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			result++
		}
	}
	return result
}

func main() {
	resultSlice := game([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(resultSlice)
}
