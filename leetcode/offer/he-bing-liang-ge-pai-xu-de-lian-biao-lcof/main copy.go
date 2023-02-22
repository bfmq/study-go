// https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	res := new(ListNode)
	node := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next, l1 = l1, l1.Next
		} else {
			node.Next, l2 = l2, l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return res.Next
}

func main() {
}
