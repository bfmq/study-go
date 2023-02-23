// https://leetcode.cn/problems/NYBBNL/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	var inorder func(*TreeNode)
	inorder = func(tn *TreeNode) {
		if tn != nil {
			inorder(tn.Left)
			vals = append(vals, tn.Val)
			inorder(tn.Right)
		}
	}
	inorder(root)

	node := &TreeNode{}
	cur := node
	for _, v := range vals {
		cur.Right = &TreeNode{Val: v}
		cur = cur.Right
	}
	return node.Right
}
