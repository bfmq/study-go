package _9_queue

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

type LinkedListQueue struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

// 入队
func (llq *LinkedListQueue) EnQueue(v interface{}) {
	node := &Node{v, nil}
	if llq.tail == nil {
		llq.tail = node
		llq.head = node
	} else {
		llq.tail.next = node
		llq.tail = node
	}
	llq.length++
}

// 出队
func (llq *LinkedListQueue) DeQueue() interface{} {
	if llq.head == nil {
		return nil
	}
	v := llq.head.val
	llq.head = llq.head.next
	llq.length--
	return v
}

// 打印
func (llq *LinkedListQueue) Print() {
	if llq.head == nil {
		fmt.Println("empty queue")
	} else {
		result := "head"
		for i := llq.head; i != nil; i = i.next {
			result += fmt.Sprintf("<-%+v", i.val)
		}
		result += "<-tail"
		fmt.Println(result)
	}
}
