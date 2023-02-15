package main

import "fmt"

func quickSort(arr []int, low, high int) []int {
	if low < high {
		left, right := low, high
		p := arr[low+(high-low)>>1]
		for left <= right {
			for arr[left] < p {
				left++
			}
			for arr[right] > p {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		quickSort(arr, low, right)
		quickSort(arr, left, high)
	}
	return arr
}

func main() {
	result := quickSort([]int{234, 65, 3, 76, 2}, 0, 4)
	fmt.Println(result)
}
