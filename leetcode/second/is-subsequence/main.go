// https://leetcode.cn/problems/is-subsequence/
package main

import "fmt"

// 两个数组下标递增比较到结束
// 如果i全部对比完毕则证明全部匹配到了
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}

func main() {
	result := isSubsequence("abc", "ahbgdc")
	fmt.Println(result)
	result = isSubsequence("axc", "ahbgdc")
	fmt.Println(result)
}
