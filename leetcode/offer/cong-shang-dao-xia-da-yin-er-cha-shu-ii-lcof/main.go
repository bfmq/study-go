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
		levelLen := len(q)
		levelNode := []int{}
		for j := 0; j < levelLen; j++ {
			node := q[0]
			q = q[1:]
			levelNode = append(levelNode, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, levelNode)
	}
	return res
}

func main() {
	result := levelOrder(&TreeNode{})
	fmt.Println(result)
}
