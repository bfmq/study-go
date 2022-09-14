// https://leetcode.cn/problems/find-smallest-letter-greater-than-target/
package main

import (
	"fmt"
	"sort"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	i := sort.Search(len(letters)-1, func(i int) bool { return letters[i] > target })
	return letters[i]
}

func main() {
	result := nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')
	fmt.Println(string(result))
	result = nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')
	fmt.Println(string(result))
	result = nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')
	fmt.Println(string(result))
}
