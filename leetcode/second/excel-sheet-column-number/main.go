// https://leetcode.cn/problems/excel-sheet-column-number/
package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	var res int
	for i, j := len(columnTitle)-1, 1; i >= 0; i--  {
		k := columnTitle[i] - 'A' + 1
		res += int(k) * j
		j *= 26
	}
	return res
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
