package main

import "fmt"

func main() {
	slice1 := []string{"I", "am", "stupid", "and", "weak"}
	slice2 := []string{"I", "am", "smart", "and", "strong"}
	for i, j := 0, 0; i < len(slice1) && j < len(slice2); i, j = i+1, j+1 {
		if slice1[i] != slice2[i] {
			slice1[i] = slice2[i]
		}
	}
	fmt.Println(slice1)
}
