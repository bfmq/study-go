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
		return nil
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		levelLen := len(q)
		levelNode := make([]int, 0, levelLen)
		for i := 0; i < levelLen; i++ {
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
	return
}

func main() {
	result := levelOrder(&TreeNode{})
	fmt.Println(result)
}
