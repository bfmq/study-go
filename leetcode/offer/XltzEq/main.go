// https://leetcode.cn/problems/XltzEq/
package main

import "strings"

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isalnum(s[l]) {
			l++
		}
		for l < r && !isalnum(s[r]) {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}
