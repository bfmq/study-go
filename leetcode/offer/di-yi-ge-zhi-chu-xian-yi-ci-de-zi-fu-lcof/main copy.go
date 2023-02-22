// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
package main

func firstUniqChar(s string) byte {
	dict := [26]int{}
	for _, c := range s {
		dict[c-'a']++
	}
	for i, c := range s {
		if dict[c-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
}
