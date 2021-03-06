package main

import (
	"fmt"
)

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

	var res = new(ListNode)
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

// func main() {
// 	l1 := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val:  4,
// 				Next: &ListNode{},
// 			},
// 		},
// 	}
// 	l2 := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 3,
// 			Next: &ListNode{
// 				Val:  4,
// 				Next: &ListNode{},
// 			},
// 		},
// 	}

// 	result := mergeTwoLists(l1, l2)
// 	fmt.Println(result)
// }

func main() {
	l1 := &ListNode{}
	fmt.Println(l1 == nil)
}
