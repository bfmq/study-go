// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		levellen := len(q)
		levelres := make([]int, 0, levellen)
		for i := 0; i < levellen; i++ {
			node := q[0]
			q = q[1:]
			levelres = append(levelres, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, levelres)
	}
	return res
}

func main() {
	result := levelOrder(&TreeNode{})
	fmt.Println(result)
}
