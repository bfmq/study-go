// https://leetcode.cn/problems/4xy4Wx/
package main

import (
	"fmt"
	"sort"
)

// 先排序
// 双指针前后算
// 出现符合条件的情况res增加j-i（都是下标不用额外计算）个符合数量，i+1继续算
// 否则j-1继续计算
func purchasePlans(nums []int, target int) int {
	var res int
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] <= target {
			res += j - i
			i++
		} else {
			j--
		}
	}
	return res % 1000000007
}

func main() {
	result := purchasePlans([]int{2, 5, 3, 5}, 6)
	fmt.Println(result)
	result = purchasePlans([]int{2, 2, 1, 9}, 10)
	fmt.Println(result)
}
