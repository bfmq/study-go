package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) (result bool) {
	if x < 0 || x > int(math.Pow(2, 31)-1) {
		// 负数及超界数不符合要求
		return false
	}
	if x == 0 {
		// 0符合要求
		return true
	}

	reserverX := func(x int) (result int) {
		// 如题7一样将数直接反转
		for x != 0 {
			result = result*10 + x%10
			x /= 10
		}
		return
	}(x)

	// 比较两数是否相等
	return x == reserverX
}

func main() {
	result := isPalindrome(121)
	fmt.Println(result)
	result = isPalindrome(-121)
	fmt.Println(result)
	result = isPalindrome(10)
	fmt.Println(result)
	result = isPalindrome(-101)
	fmt.Println(result)
}
