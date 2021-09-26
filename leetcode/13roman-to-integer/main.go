package main

import (
	"fmt"
)

func romanToInt(s string) (result int) {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sLength := len(s)
	for i := 0; i < len(s)-1; i++ {
		currStr := fmt.Sprintf("%c", s[i])
		nextStr := fmt.Sprintf("%c", s[i+1])
		if romanMap[currStr] < romanMap[nextStr] {
			result -= romanMap[currStr]
		} else {
			result += romanMap[currStr]
		}
	}

	lastStr := fmt.Sprintf("%c", s[sLength-1])
	result += romanMap[lastStr]

	return
}

func main() {
	result := romanToInt("LVIII")
	fmt.Println(result)
}
