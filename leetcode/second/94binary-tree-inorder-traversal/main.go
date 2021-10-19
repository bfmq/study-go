package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (result []int) {
	// 定义中序变量函数
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 左递归
		inorder(node.Left)
		// 根添加
		result = append(result, node.Val)
		// 右递归
		inorder(node.Right)
	}
	inorder(root)
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
	result := inorderTraversal(treeNode1)
	fmt.Println(result)

	treeNode2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	result = inorderTraversal(treeNode2)
	fmt.Println(result)
}
