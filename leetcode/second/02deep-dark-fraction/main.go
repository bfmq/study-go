package main

import "fmt"

func fraction(cont []int) []int {
	if len(cont) == 1 {
		return []int{cont[0], 1}
	} else {
		last := fraction(cont[1:])
		return []int{cont[0]*last[0] + last[1], last[0]}
	}
}

func main() {
	result := fraction([]int{3, 2, 1, 2})
	fmt.Println(result)
	result = fraction([]int{0, 0, 3})
	fmt.Println(result)
}
