// https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	pre := head
	for pre.Next != nil && pre.Next.Val != val {
		pre = pre.Next
	}
	if pre.Next != nil {
		pre.Next = pre.Next.Next
	}
	return head
}

func main() {
}
