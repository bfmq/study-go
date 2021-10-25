package main

import (
	"math"
)

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

func maxPathSum(root *TreeNode) int {
	sumMax := math.MinInt64
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(helper(node.Left), 0)
		right := max(helper(node.Right), 0)

		sumMax = max(sumMax, node.Val+left+right)
		return node.Val + max(left, right)
	}
	helper(root)
	return sumMax
}

func main() {
}
