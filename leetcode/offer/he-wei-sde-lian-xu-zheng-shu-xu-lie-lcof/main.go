// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
package main

func findContinuousSequence(target int) (res [][]int) {
	i, j, s := 1, 2, 3
	for i < j {
		if s == target {
			m := make([]int, 0, j-i)
			for z := i; z < j; z++ {
				m = append(m, z)
			}
			res = append(res, m)
		} else if s < target {
			s += j
			j++
		} else {
			s -= i
			i++
		}
	}
	return res
}

func main() {
}
