// https://leetcode.cn/problems/2vYnGI/
package main

import (
	"fmt"
	"sort"
)

// 先排序
// 按staple正算，drinks反算；
// 第i个staple情况下，第j个drinks满足条件，则j前面所有drinks都满足，结果增加j+1（因为j是下标，真实个数+1），此i计算完毕+1
// 否则j-1继续计算
func breakfastNumber(staple []int, drinks []int, x int) int {
	var res int
	sort.Ints(staple)
	sort.Ints(drinks)
	for i, j := 0, len(drinks)-1; i < len(staple) && j >= 0; {
		if staple[i]+drinks[j] <= x {
			res += j + 1
			i++
		} else {
			j--
		}
	}
	return res % 1000000007
}

func main() {
	result := breakfastNumber([]int{10, 20, 5}, []int{5, 5, 2}, 15)
	fmt.Println(result)
	result = breakfastNumber([]int{2, 1, 1}, []int{8, 9, 5, 1}, 9)
	fmt.Println(result)
}
