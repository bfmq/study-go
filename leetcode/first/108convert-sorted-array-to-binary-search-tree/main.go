package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return makeNode(nums, 0, len(nums)-1)
}

func makeNode(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = makeNode(nums, left, mid-1)
	root.Right = makeNode(nums, mid+1, right)
	return root
}

func main() {
	result := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(result)
}
