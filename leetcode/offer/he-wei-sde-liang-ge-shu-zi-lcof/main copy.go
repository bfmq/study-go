// https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/
package main

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		s := nums[l] + nums[r]
		if s < target {
			l++
		} else if s > target {
			r--
		} else {
			return []int{nums[l], nums[r]}
		}
	}
	return nil
}

func main() {
}
