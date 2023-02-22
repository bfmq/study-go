// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) (res []int) {
	if head == nil {
		return nil
	}

	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	for pre != nil {
		res = append(res, pre.Val)
		pre = pre.Next
	}
	return
}

func main() {
}
