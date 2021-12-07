package main

import "fmt"

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArea(height []int) (result int) {
	i := 0
	j := len(height) - 1
	for i < j {
		area := (j - i) * min(height[i], height[j])
		result = max(area, result)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return
}

func main() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(result)
	result = maxArea([]int{1, 1})
	fmt.Println(result)
	result = maxArea([]int{1, 2, 1})
	fmt.Println(result)
}
