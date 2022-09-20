// https://leetcode.cn/problems/longest-common-prefix/
package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) (result string) {
	strsLength := len(strs)
	// 因为题目说了不会出现空strs情况
	// 长度为1直接返回
	if strsLength == 1 {
		return strs[0]
	}

	// 对切片排序，会自动按acsii排序
	sort.Strings(strs)
	// 只需要比较第一个str跟最后str的差异性
	mi, ma := strs[0], strs[strsLength-1]
	for i := 0; i < len(mi); i++ {
		if mi[i] != ma[i] {
			// 出现不等时，则结束
			return mi[:i]
		}
	}
	// 否则说明最后str完全包含了首个str，因此返回首个str
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
