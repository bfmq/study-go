package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 入参检测
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 当前值变更
	root1.Val += root2.Val
	// 左节点递归变更
	root1.Left = mergeTrees(root1.Left, root2.Left)
	// 右节点递归变更
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func main() {
}
