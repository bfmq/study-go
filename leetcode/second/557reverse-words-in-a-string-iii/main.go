package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// 切分成单个字符串
	sSlice := strings.Split(s, " ")
	// 定义返回值切片
	res := make([]string, 0, len(sSlice))
	for _, ss := range sSlice {
		// 单个字符串切分分单个字符
		sL := strings.Split(ss, "")
		// 下标前后互换
		for i, j := 0, len(sL)-1; i < j; i, j = i+1, j-1 {
			sL[i], sL[j] = sL[j], sL[i]
		}
		// 返回值切片添加本字符串反转结果
		res = append(res, strings.Join(sL, ""))
	}
	// 拼接返回值切片
	return strings.Join(res, " ")
}

func main() {
	result := reverseWords("Let's take LeetCode contest")
	fmt.Println(result)
}
