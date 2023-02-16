// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/
package main

import "math"

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(maxDepth(root.Left))-float64(maxDepth(root.Right))) <= 1 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)
}

func main() {
}
