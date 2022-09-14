// https://leetcode.cn/problems/first-bad-version/
package main

import "fmt"

func isBadVersion(version int) bool {
	if version < 4 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	// 搜索数字，不是下标，因此是1,n
	l, r := 1, n
	for l < r {
		m := l + (r-l)>>1
		if isBadVersion(m) {
			// 是正确版本则r变更
			r = m
		} else {
			// 是错误版本则l变更
			l = m + 1
		}
	}
	return r
}

func main() {
	result := firstBadVersion(5)
	fmt.Println(result)
	// result = firstBadVersion(1)
	// fmt.Println(result)
}
