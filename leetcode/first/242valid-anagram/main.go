package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sl [26]int
	var tl [26]int
	for _, x := range s {
		sl[x-'a'] += 1
	}
	for _, x := range t {
		tl[x-'a'] += 1
	}
	return sl == tl
}

func main() {
	result := isAnagram("anagram", "nagaram")
	fmt.Println(result)
	result = isAnagram("rat", "car")
	fmt.Println(result)
}
