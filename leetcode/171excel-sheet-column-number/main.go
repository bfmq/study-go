package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) (result int) {
	for i, n := len(columnTitle)-1, 1; i >= 0; i, n = i-1, n*26 {
		column := columnTitle[i] - 'A' + 1
		result += int(column) * n
	}
	return
}

func main() {
	result := titleToNumber("A")
	fmt.Println(result)
	result = titleToNumber("AB")
	fmt.Println(result)
	result = titleToNumber("ZY")
	fmt.Println(result)
	result = titleToNumber("FXSHRXW")
	fmt.Println(result)
}
