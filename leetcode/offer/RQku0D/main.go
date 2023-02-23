// https://leetcode.cn/problems/RQku0D/
package main

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			f1, f2 := true, true
			for i, j := l, r-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					f1 = false
					break
				}
			}
			for i, j := l+1, r; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					f2 = false
					break
				}
			}
			return f1 || f2
		}
	}
	return true
}
