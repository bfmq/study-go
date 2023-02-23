// https://leetcode.cn/problems/aMhZSa/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	n := len(vals)
	for i, v := range vals[:n>>1] {
		if v != vals[n-i-1] {
			return false
		}
	}
	return true
}
