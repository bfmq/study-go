// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var res int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			res = node.Val
		}
		dfs(node.Left)
	}
	dfs(root)
	return res
}

func main() {
}
