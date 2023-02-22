// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var check func(*TreeNode, *TreeNode) bool
	check = func(tn1, tn2 *TreeNode) bool {
		if tn1 == nil && tn2 == nil {
			return true
		}
		if tn1 == nil || tn2 == nil {
			return false
		}
		return tn1.Val == tn2.Val && check(tn1.Left, tn2.Right) && check(tn1.Right, tn2.Left)
	}
	return check(root, root)
}

func main() {
}
