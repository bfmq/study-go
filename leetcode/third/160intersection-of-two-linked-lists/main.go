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
		if heada == nil {
			heada = headB
		} else {
			heada = heada.Next
		}
		if headb == nil {
			headb = headA
		} else {
			headb = headb.Next
		}
	}
	return heada
}

func main() {

}
