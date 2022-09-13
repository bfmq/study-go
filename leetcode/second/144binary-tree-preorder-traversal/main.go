package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (result []int) {
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return
}

func main() {
	treeNode1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	result := preorderTraversal(treeNode1)
	fmt.Println(result)

	treeNode2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	result = preorderTraversal(treeNode2)
	fmt.Println(result)
}
