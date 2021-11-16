package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	// needle为空或两串相等直接反0
	if needle == "" || haystack == needle {
		return 0
	}

	// needle比haystack长反-1
	h, n := len(haystack), len(needle)
	if n > h {
		return -1
	}

	// 比较到少于n的长度即可
	for i := 0; i < (h - n + 1); i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	// 找不到反-1
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
