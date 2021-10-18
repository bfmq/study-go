package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) (result bool) {
	if x < int(math.Pow(-2, 31)) || x > int(math.Pow(2, 31))-1 || x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	reverseX := func(x int) (r int) {
		for {
			if x != 0 {
				r = r*10 + x%10
				x /= 10
			} else {
				break
			}
		}
		return
	}(x)

	return x == reverseX
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
