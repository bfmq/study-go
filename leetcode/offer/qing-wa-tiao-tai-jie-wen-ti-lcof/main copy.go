// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
package main

func numWays(n int) int {
	const mod int = 1e9 + 7
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}

func main() {
}
