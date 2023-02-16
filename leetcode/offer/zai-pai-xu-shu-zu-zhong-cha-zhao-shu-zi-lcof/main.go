// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
package main

func search(nums []int, target int) int {
	var res int
	for _, n := range nums {
		if n == target {
			res++
		}
	}
	return res
}

func main() {
}
