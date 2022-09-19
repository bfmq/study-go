// https://leetcode.cn/problems/path-sum/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 结束了比较最终值
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	// 左右有一路径满足即可
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
}
