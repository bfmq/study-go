package main

import "fmt"

func quickSort(arr []int, low, high int) []int {
	if low < high {
		l, r := low, high
		p := arr[l+(r-l)>>1]
		for l <= r {
			for arr[l] < p {
				l++
			}
			for arr[r] > p {
				r--
			}
			if l < r {
				arr[l], arr[r] = arr[r], arr[l]
				l++
				r--
			}
		}
		quickSort(arr, low, r)
		quickSort(arr, l, high)
	}
	return arr
}

func main() {
	result := quickSort([]int{234, 65, 3, 76, 2}, 0, 4)
	fmt.Println(result)
}
