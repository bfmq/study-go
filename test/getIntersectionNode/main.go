package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	heada, headb := headA, headB
	for heada != headb {
		if heada != nil {
			heada = heada.Next
		} else {
			heada = headB
		}
		if headb != nil {
			headb = headb.Next
		} else {
			headb = headA
		}
	}
	return heada
}

func main() {

}
