package main

import "fmt"

func game(guess []int, answer []int) int {
	// 3次直接循环对比即可
	var res int
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			res++
		}
	}
	return res
}

func main() {
	resultSlice := game([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(resultSlice)
}
