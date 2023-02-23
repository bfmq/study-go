// https://leetcode.cn/problems/dKk3P7/
package main

import "sort"

func isAnagram(s string, t string) bool {
	if s == t {
		return false
	}
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}
