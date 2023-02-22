// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
package main

func findContinuousSequence(target int) (res [][]int) {
	x, y, z := 1, 2, 3
	for x < y {
		if z == target {
			r := make([]int, 0, y-x)
			for i := x; i < y; i++ {
				r = append(r, i)
			}
			res = append(res, r)
		} else if z < target {
			z += y
			y++
		} else {
			z -= x
			x++
		}
	}
	return
}

func main() {
}
