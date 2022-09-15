// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
package main

import "fmt"

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 数据全部已知是前提
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0
	// 获取可以拿到的最低价
	// 用今日价格减最低价获得最大利益
	for _, price := range prices {
		maxProfit = max(maxProfit, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return maxProfit
}

func main() {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(result)
	result = maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Println(result)
}
