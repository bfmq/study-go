// https://leetcode.cn/problems/sZ59z6/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左右递归即可
func numColor(root *TreeNode) int {
	colorMap := make(map[int]bool)
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn != nil {
			colorMap[tn.Val] = false
			dfs(tn.Left)
			dfs(tn.Right)
		}
	}
	dfs(root)
	return len(colorMap)
}

func main() {

}
