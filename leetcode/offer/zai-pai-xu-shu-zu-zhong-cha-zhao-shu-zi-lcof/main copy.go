// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
package main

func search(nums []int, target int) (res int) {
	for _, v := range nums {
		if v == target {
			res++
		}
	}
	return
}

func main() {
}
