package _7_linkedlist

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head *ListNode
}

// 反转
func (ll *LinkedList) Reverse() {
	if ll.head == nil || ll.head.next == nil || ll.head.next.next == nil {
		return
	}

	var pre *ListNode = nil
	cur := ll.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	ll.head.next = pre
}

// 环检测
func (ll *LinkedList) HasCycle() bool {
	if ll.head != nil {
		slow := ll.head
		fast := ll.head
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if fast == slow {
				return true
			}
		}
	}
	return false
}

