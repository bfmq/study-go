package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (result []int) {
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return
}

func main() {
	treeNode1 := &TreeNode{
		Val:  1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
			},
		},
	}
	result := inorderTraversal(treeNode1)
	fmt.Println(result)

	treeNode2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
		},
	}
	result = inorderTraversal(treeNode2)
	fmt.Println(result)
}
