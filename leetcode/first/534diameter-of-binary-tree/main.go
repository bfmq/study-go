package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func diameterOfBinaryTree(root *TreeNode) int {
	result := 1
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		L := depth(node.Left)
		R := depth(node.Right)
		result = max(L+R+1, result)
		return max(L, R) + 1
	}

	_ = depth(root)
	return result - 1
}

func main() {
}
