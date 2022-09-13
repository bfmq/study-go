// https://leetcode.cn/problems/qi-wang-ge-shu-tong-ji/
package main

import "fmt"

func expectNumber(scores []int) int {
	// 本质就是求set长度
	hash := make(map[int]bool)
	for _, score := range scores {
		hash[score] = false
	}
	return len(hash)
}

func main() {
	result := expectNumber([]int{1, 2, 3})
	fmt.Println(result)
	result = expectNumber([]int{1, 1})
	fmt.Println(result)
	result = expectNumber([]int{1, 1, 2})
	fmt.Println(result)
}
