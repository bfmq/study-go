// https://leetcode.cn/problems/valid-parentheses/
package main

import (
	"fmt"
)

func isValid(s string) bool {
	n := len(s)
	// 如果s长度直接为奇数，那么肯定无法对称
	if n%2 == 1 {
		return false
	}

	// 基础比对map
	baseMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 创建一个栈（伪）
	stack := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if _, ok := baseMap[s[i]]; ok {
			// 如果查到了
			// 1.当前栈是否为空，为空没得对称
			// 2.比对是否为对称括号
			// 3.出栈
			if len(stack) == 0 || stack[len(stack)-1] != baseMap[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 如果查不到，则入栈
			stack = append(stack, s[i])
		}
	}

	// 验证最终栈是否为空
	return len(stack) == 0
}

func main() {
	result := isValid("[][][]")
	fmt.Println(result)
	result = isValid("()[]{}")
	fmt.Println(result)
	result = isValid("(]")
	fmt.Println(result)
	result = isValid("([)]")
	fmt.Println(result)
	result = isValid("{[]}")
	fmt.Println(result)
}
