// https://leetcode.cn/problems/pascals-triangle/
package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
		// 0跟末尾固定为1
		res[i][0] = 1
		res[i][i] = 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

func main() {
	result := generate(5)
	fmt.Println(result)
	result = generate(1)
	fmt.Println(result)
}
