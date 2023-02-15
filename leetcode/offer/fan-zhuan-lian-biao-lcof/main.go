// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	latter, former := head, head
	for i := 0; i < k; i++ {
		former = former.Next
	}
	for former != nil {
		latter = latter.Next
		former = former.Next
	}
	return latter
}

func main() {
}
