// https://leetcode.cn/problems/check-if-n-and-its-double-exist/
package main

import "fmt"

func checkIfExist(arr []int) bool {
	// 收集下标
	hash := make(map[int]int, len(arr))
	for i, v := range arr {
		hash[v] = i
	}

	// *2后对比下标不同则说明存在2倍
	for i, v := range arr {
		if k, ok := hash[v<<1]; i != k && ok {
			return true
		}
	}
	return false
}

func main() {
	result := checkIfExist([]int{10, 2, 5, 3})
	fmt.Println(result)
	result = checkIfExist([]int{7, 1, 14, 11})
	fmt.Println(result)
	result = checkIfExist([]int{3, 1, 7, 11})
	fmt.Println(result)
}
