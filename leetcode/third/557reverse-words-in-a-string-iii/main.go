package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sLs := strings.Split(s, " ")
	result := make([]string, 0, len(sLs))
	for _, sL := range sLs {
		s := strings.Split(sL, "")
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		result = append(result, strings.Join(s, ""))
	}
	return strings.Join(result, " ")
}

func main() {
	result := reverseWords("Let's take LeetCode contest")
	fmt.Println(result)
}
