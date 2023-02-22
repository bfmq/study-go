// https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

func isStraight(nums []int) bool {
	mi, ma := 14, 0
	isExist := make(map[int]bool, len(nums))
	for _, num := range nums {
		if num == 0 {
			continue
		}
		mi = min(mi, num)
		ma = max(ma, num)
		if _, ok := isExist[num]; ok {
			return false
		} else {
			isExist[num] = true
		}
	}
	return ma-mi < 5
}

func main() {
}
