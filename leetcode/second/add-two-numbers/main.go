// https://leetcode.cn/problems/add-two-numbers/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		// 取l1跟l2当前值，否则设为0
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		// 本次总和=上次余数 + n1 + n2
		sum := carry + n1 + n2
		// 计算商数跟余数
		sum, carry = sum%10, sum/10
		if head == nil {
			// 初始化返回值
			head = &ListNode{Val: sum}
			tail = head
		} else {
			// 追加后续值
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	// 处理最后一次处理的余数
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
}
