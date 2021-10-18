package main

import "fmt"

func isAnagram(s string, t string) bool {
	// 长度不等直接错
	if len(s) != len(t) {
		return false
	}

	// 遍历s进行计数加法
	stMap := make(map[rune]int)
	for _, c := range s {
		stMap[c]++
	}

	// 遍历t进行计数减法，出现小于0则已经不等了
	for _, c := range t {
		stMap[c]--
		if stMap[c] < 0 {
			return false
		}
	}
	return true
}

func main() {
	result := isAnagram("anagram", "nagaram")
	fmt.Println(result)
	result = isAnagram("rat", "car")
	fmt.Println(result)
}
