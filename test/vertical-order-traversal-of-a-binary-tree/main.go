package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	row, col, val int
}

func verticalTraversal(root *TreeNode) [][]int {
	var nodes []*NodeInfo
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		nodes = append(nodes, &NodeInfo{row: row, col: col, val: node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		x, y := nodes[i], nodes[j]
		if x.col != y.col {
			return x.col < y.col
		}
		if x.row != y.row {
			return x.row < y.row
		}
		return x.val < y.val
	})

	var res [][]int
	cur := []int{nodes[0].val}
	preCol := nodes[0].col
	for i := 1; i < len(nodes); i++ {
		node := nodes[i]
		col := node.col
		val := node.val
		if col == preCol {
			cur = append(cur, val)
		} else {
			res = append(res, cur)
			cur = []int{val}
			preCol = node.col
		}
	}
	res = append(res, cur)
	return res
}

func main() {

}
