// https://leetcode.cn/problems/univalued-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return root == nil || (root.Left == nil || root.Val == root.Left.Val && isUnivalTree(root.Left)) &&
		(root.Right == nil || root.Val == root.Right.Val && isUnivalTree(root.Right))

}

func main() {
}
