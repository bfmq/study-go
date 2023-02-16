// https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ha, hb := headA, headB
	for ha != hb {
		if ha != nil {
			ha = ha.Next
		} else {
			ha = headB
		}
		if hb != nil {
			hb = hb.Next
		} else {
			hb = headA
		}
	}
	return ha
}

func main() {
}
