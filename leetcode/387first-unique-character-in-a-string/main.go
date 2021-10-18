package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	sArray := [26]int{}
	for _, ch := range s {
		sArray[ch-'a']++
	}

	for i, ch := range s {
		if sArray[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	result := firstUniqChar("leetcode")
	fmt.Println(result)
	result = firstUniqChar("loveleetcode")
	fmt.Println(result)
}
