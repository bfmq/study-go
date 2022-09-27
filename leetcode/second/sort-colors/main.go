// https://leetcode.cn/problems/sort-colors/
package main

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i <= r; i++ {
		for ; i <= r && nums[i] == 2; r-- {
			nums[i], nums[r] = nums[r], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
}

func main() {
}
