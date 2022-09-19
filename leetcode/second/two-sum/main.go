// https://leetcode.cn/problems/two-sum/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// key为值，value为下标
	nMap := make(map[int]int, len(nums))

	// 遍历循环
	for i, num := range nums {
		if j, ok := nMap[target-num]; ok {
			// target-num在字典里，那么取当前遍历到的i跟字典里的对应下标
			return []int{i, j}
		} else {
			// 否则说明遍历至此时前面的数没有满足相加等于target的
			nMap[num] = i
		}
	}

	// 题目声明一定会有结果，这里其实永远不会返回
	return nil
}

func main() {
	resultSlice := twoSum([]int{2, 7, 11, 15}, 26)
	fmt.Println(resultSlice)
}
