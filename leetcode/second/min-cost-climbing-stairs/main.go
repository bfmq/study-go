// https://leetcode.cn/problems/min-cost-climbing-stairs/
package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	n := len(cost)
	// 按题目描述最后一个台阶也算
	// 本次的台阶可以从前一步或两步上来
	// 该台阶最小价值=该台阶本身价值+前一/二可的最小价值
	for i := 2; i < n; i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[n-1], cost[n-2])
}

func main() {
	result := minCostClimbingStairs([]int{10, 15, 20})
	fmt.Println(result)
	result = minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	fmt.Println(result)
}
