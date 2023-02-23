// https://leetcode.cn/problems/w3tCBm/
package main

func ones(n int) (res int) {
	for ; n > 0; n &= n - 1 {
		res++
	}
	return
}

func countBits(n int) (res []int) {
	for i := 0; i <= n; i++ {
		res = append(res, ones(i))
	}
	return res
}
