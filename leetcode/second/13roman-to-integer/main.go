package main

import (
	"fmt"
)

// 可以理解成是一道找规律题
func romanToInt(s string) (result int) {
	// 定义基础map
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// 循环下标
	for i := range s {
		// 取下标值
		v := romanMap[s[i]]
		if i < len(s)-1 && v < romanMap[s[i+1]] {
			// 如果下标值比next下标值小则减
			// 注意i不能越界，最后一个值必须加
			result -= v
		} else {
			// 反之则加
			result += v
		}
	}
	return
}

func main() {
	result := romanToInt("LVIII")
	fmt.Println(result)
}
