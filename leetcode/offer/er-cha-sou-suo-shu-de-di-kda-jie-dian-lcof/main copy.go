// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) (res int) {
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(root.Right)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			res = tn.Val
		}
		dfs(root.Left)
	}
	dfs(root)
	return
}

func main() {
}
