package main

import (
	"fmt"
	"math"
)

func reverse(x int) (result int) {
	for {
		if x != 0 {
			result = result*10 + x%10
			x /= 10
		} else {
			if result > int(math.Pow(2, 31))-1 || result < int(math.Pow(-2, 31)) {
				result = 0
			}
			break
		}
	}

	return
}

func main() {
	result := reverse(-123)
	fmt.Println(result)
}
