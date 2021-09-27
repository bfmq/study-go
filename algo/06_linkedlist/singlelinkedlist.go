package _6_linkedlist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

// 初始化节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		value: v,
		next:  nil,
	}
}

// 获取节点下一指向
func (ln *ListNode) GetNext() *ListNode {
	return ln.next
}

// 获取节点值
func (ln *ListNode) GetValue() interface{} {
	return ln.value
}

// 初始化链表
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   NewListNode(0),
		length: 0,
	}
}

// 在某节点后插入新节点
func (ll *LinkedList) InsertAfter(ln *ListNode, v interface{}) bool {
	if ln == nil {
		return false
	}

	newNode := NewListNode(v)
	// 逻辑顺序不可变
	newNode.next = ln.next
	ln.next = newNode
	ll.length++
	return true
}

// 头部插入节点
func (ll *LinkedList) InsertToHead(v interface{}) bool {
	return ll.InsertAfter(ll.head, v)
}

// 尾部插入节点
func (ll *LinkedList) InsertToTail(v interface{}) bool {
	// 循环到最后才能插入新节点
	cur := ll.head
	for cur != nil {
		cur = cur.next
	}
	return ll.InsertAfter(cur, v)
}

// 在某节点前插入新节点
func (ll *LinkedList) InsertBefore(ln *ListNode, v interface{}) bool {
	if ln == nil || ln == ll.head {
		return false
	}

	cur := ll.head.next
	pre := ll.head

	for cur != nil {
		if cur == ln {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	ll.length++
	return true
}

// 通过索引查找节点
func (ll *LinkedList) FindByIndex(index uint) *ListNode {
	if index > ll.length {
		return nil
	}

	cur := ll.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除传入的节点
func (ll *LinkedList) DeleteListNode(ln *ListNode) bool {
	if ln == nil || ll.length == 0 {
		return false
	}

	cur := ll.head.next
	pre := ll.head

	for cur != nil {
		if cur == ln {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}
	pre.next = ln.next
	ln = nil
	ll.length--
	return true
}

// 打印
func (ll *LinkedList) Print() {
	var s string
	for ln := ll.head.next; ln != nil; ln = ln.next {
		s += fmt.Sprintf("->%+v", ln.value)
	}
	fmt.Println(s)
}
