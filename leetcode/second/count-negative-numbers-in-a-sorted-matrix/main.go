// https://leetcode.cn/problems/count-negative-numbers-in-a-sorted-matrix/
package main

func countNegatives(grid [][]int) int {
	var res int
	for _, i := range grid {
		for _, j := range i {
			if j < 0 {
				res++
			}
		}
	}
	return res
}

func main() {
}
