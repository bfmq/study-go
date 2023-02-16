// https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/
package main

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1

	for i < j {
		s := nums[i] + nums[j]
		if s > target {
			j--
		} else if s < target {
			i++
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return nil
}

func main() {
}
