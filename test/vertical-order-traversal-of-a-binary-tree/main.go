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
	var allNode []*NodeInfo
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		allNode = append(allNode, &NodeInfo{row: row, col: col, val: node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	sort.Slice(allNode, func(i, j int) bool {
		if allNode[i].col != allNode[j].col {
			return allNode[i].col < allNode[j].col
		}
		if allNode[i].row != allNode[j].row {
			return allNode[i].row < allNode[j].row
		}
		return allNode[i].val < allNode[j].val
	})

	var res [][]int
	cur := []int{allNode[0].val}
	preCol := allNode[0].col
	for i := 1; i < len(allNode); i++ {
		node := allNode[i]
		col := node.col
		val := node.val
		if preCol == col {
			cur = append(cur, val)
		} else {
			res = append(res, cur)
			preCol = col
			cur = []int{val}
		}
	}
	res = append(res, cur)
	return res
}

func main() {

}
