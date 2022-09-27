// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 计算长度
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	// 到n前正常追加
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	// 单独跳过n
	cur.Next = cur.Next.Next
	return dummy.Next
}

func main() {
}
