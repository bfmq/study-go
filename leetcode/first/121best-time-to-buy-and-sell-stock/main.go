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

func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	minNow := prices[0]
	maxNow := 0

	for _, price := range prices {
		minNow = min(price, minNow)
		maxNow = max(maxNow, price-minNow)
	}
	return maxNow
}

func main() {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(result)
	result = maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Println(result)
}
