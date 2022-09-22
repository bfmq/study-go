// https://leetcode.cn/problems/3sum/
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			// nums[i] > 0说明后面的都不满足相加为0了
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			// 之前算过的数字跳过
			continue
		}
		l := i + 1 // 下一个下标
		r := n - 1 // 最后一位下标
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				// 满足添加
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					// 跳过重复数字
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					// 跳过重复数字
					r--
				}
				// 两边下标夹逼
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] > 0 {
				// 大于0需要减少大数
				r--
			} else {
				// 小于0需要增加小数
				l++
			}
		}
	}
	return res
}

func main() {
	resultSlice := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(resultSlice)
	resultSlice = threeSum([]int{0, 1, 1})
	fmt.Println(resultSlice)
	resultSlice = threeSum([]int{0, 0, 0})
	fmt.Println(resultSlice)
}
