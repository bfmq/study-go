// https://leetcode.cn/problems/nGK0Fy/
package main

import "fmt"

// 目标结果是x+y
// 出现一个"A"，有x+y=(2x+y)+y=2x+2y
// 出现一个"B"，有x+y=x+(2y+x)=2x+2y
// 所以每出现一个A/B，都使x+y的值翻倍
// 因此结果是2的len(s)次方
// golang中^是位异或
// 因此使用位运算
// a := 1 << 3 // 2的3次方*1
// b := 1 << 6 // 2的6次方*1
// c := 4 << 2 // 2的2次方*4
// d := 4 << 3 // 2的3次方*4
// a := 16 >> 3 // 16除以2的3次方
func calculate(s string) int {
	return 1 << len(s)
}

func main() {
	result := calculate("AB")
	fmt.Println(result)
}
