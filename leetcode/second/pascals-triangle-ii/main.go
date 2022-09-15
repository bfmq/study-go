// https://leetcode.cn/problems/pascals-triangle-ii/
package main

import "fmt"

func getRow(rowIndex int) []int {
	// 定义两个切片用于互相生成即可
	var pre, cur []int
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j] + pre[j-1]
		}
		pre = cur
	}
	return pre
}

func main() {
	result := getRow(5)
	fmt.Println(result)
	result = getRow(1)
	fmt.Println(result)
}
