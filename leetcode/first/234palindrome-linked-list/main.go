package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	l := make([]int, 0, 10)
	for node := head; node != nil; node = node.Next {
		l = append(l, node.Val)
	}
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		if l[i] != l[j] {
			return false
		}
	}
	return true
}

func main() {
	result := isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	})
	fmt.Println(result)
	result = isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	})
	fmt.Println(result)
}
