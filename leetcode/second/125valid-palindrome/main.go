package main

import (
	"fmt"
	"strings"
)

func checkChar(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		for i < j && !checkChar(s[i]) {
			i = i + 1
		}
		for i < j && !checkChar(s[j]) {
			j = j - 1
		}
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	result := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(result)
	result = isPalindrome("race a car")
	fmt.Println(result)
}
