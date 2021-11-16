package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 入参检测
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 创建一个新节点
	node := new(ListNode)
	// 赋值给变量节点，这样最终返回元节点node即可，如果只有一个节点无法直接返回node.Next
	res := node
	// l1 l2都有值则可一直比较
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			// l1大则添加l2的
			res.Next, l2 = l2, l2.Next
		} else {
			// l2大则添加l1的
			res.Next, l1 = l1, l1.Next
		}
		// 每次添加完，res需要指向res.Next
		res = res.Next
	}

	// 结束验证
	if l1 != nil {
		res.Next = l1
	}
	if l2 != nil {
		res.Next = l2
	}

	// 直接返回没有污染过的node即可，res此时已经是在靠后位置了
	return node.Next
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
