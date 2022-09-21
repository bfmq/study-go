// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]bool{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && !m[s[rk+1]] {
			// 从每个字符慢慢向后递增测试是否重复
			m[s[rk+1]] = true
			rk++
		}
		// 比较该s[i]可达到的最大长度跟之前达到的最大长度
		ans = max(ans, rk-i+1)
	}
	return ans
}

func main() {
}
