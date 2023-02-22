// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
package main

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func main() {
}
