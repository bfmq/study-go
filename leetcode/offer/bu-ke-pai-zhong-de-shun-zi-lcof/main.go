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
	isExist := make(map[int]bool, len(nums))
	mi, ma := 14, 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		ma = max(ma, num)
		mi = min(mi, num)
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
