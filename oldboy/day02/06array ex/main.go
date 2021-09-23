package main

import "fmt"

func main() {
	sums := [...]int{1, 3, 5, 7, 8}
	su := 0
	for _, v := range sums {
		su += v
	}
	fmt.Println(su)

	for i, v := range sums {
		for j := i + 1; j < len(sums); j++ {
			if v+sums[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}

}
