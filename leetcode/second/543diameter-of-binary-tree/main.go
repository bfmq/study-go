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
	var res int
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 左节点可提供最大深度
		l := depth(node.Left)
		// 右节点可提供最大深度
		r := depth(node.Right)
		// 该节点本身形成的树与当前最大值比较
		res = max(res, l+r+1)
		// 返回该节点可为父节点提供的最大深度
		// 因为要与父节点形成树，因此只能取左/右分支，否则无法满足题目要求
		return max(l, r) + 1
	}

	depth(root)
	return res - 1
}

func main() {
}
