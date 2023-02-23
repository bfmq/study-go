// https://leetcode.cn/problems/jJ0w9p/
package main

func mySqrt(x int) (res int) {
	l, r := 0, x
	for l <= r {
		m := l + (r-l)>>1
		if m*m <= x {
			res = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return
}
