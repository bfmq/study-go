package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if haystack == needle || needle == "" {
		return 0
	}

	h, n := len(haystack), len(needle)
	if n > h {
		return -1
	}

	for i := 0; i < (h - n + 1); i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}

func main() {
	result := strStr("hello", "ll")
	fmt.Println(result)
	result = strStr("aaaaa", "bba")
	fmt.Println(result)
	result = strStr("", "")
	fmt.Println(result)
	result = strStr("aaaaa", "aaaaaaaaaaaaaa")
	fmt.Println(result)
	result = strStr("abc", "c")
	fmt.Println(result)
}
