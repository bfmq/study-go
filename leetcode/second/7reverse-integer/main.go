package main

import (
	"fmt"
	"math"
)

func reverse(x int) (result int) {
	for {
		// x!=0表示x还没被处理完
		if x != 0 {
			// 1.result首次是0不影响乘法
			// 2.循环计算x的余数
			// 3.每次乘十并累加新的
			// 4.x除掉10
			result = result*10 + x%10
			x /= 10
		} else {
			if result < int(math.Pow(-2, 31)) || result > int(math.Pow(2, 31))-1 {
				// 结果反而可能会出现越界情况
				result = 0
			}
			return
		}
	}
}

func main() {
	result := reverse(-123)
	fmt.Println(result)
}
