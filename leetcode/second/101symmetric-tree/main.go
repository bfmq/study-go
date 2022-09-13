package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var checkNode func(leftNode, rightNode *TreeNode) bool
	checkNode = func(leftNode, rightNode *TreeNode) bool {
		if leftNode == nil && rightNode == nil {
			return true
		}
		if leftNode == nil || rightNode == nil {
			return false
		}
		if leftNode.Val != rightNode.Val {
			return false
		}
		return checkNode(leftNode.Left, rightNode.Right) && checkNode(rightNode.Left, leftNode.Right)
	}
	return checkNode(root.Left, root.Right)
}

func main() {
	treeNode1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	result := isSymmetric(treeNode1)
	fmt.Println(result)

	treeNode2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	result = isSymmetric(treeNode2)
	fmt.Println(result)
}
