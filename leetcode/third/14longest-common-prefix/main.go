package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) (result string) {
	strsLength := len(strs)
	if strsLength == 0 {
		return
	}
	if strsLength == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	mi, ma := strs[0], strs[len(strs)-1]
	for i := 0; i < len(mi); i++ {
		if mi[i] != ma[i] {
			return mi[:i]
		}
	}

	return mi
}

func main() {
	result := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Println(result)
	result = longestCommonPrefix([]string{"ab", "a"})
	fmt.Println(result)
	result = longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(result)
}
