package _9_queue

import "fmt"

type CircularQueue struct {
	queue    []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		queue:    make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// 判断队列是否满
func (cq *CircularQueue) IsFull() bool {
	if cq.head == (cq.tail+1)%cq.capacity {
		return true
	}
	return false
}

// 判断队列是否空
func (cq *CircularQueue) IsEmpty() bool {
	if cq.head == cq.tail {
		return true
	}
	return false
}

// 入队
func (cq *CircularQueue) EnQueue(v interface{}) bool {
	if cq.IsFull() {
		return false
	}
	cq.queue[cq.tail] = v
	cq.tail = (cq.tail + 1) % cq.capacity
	return true
}

// 出队
func (cq *CircularQueue) DeQueue() interface{} {
	if cq.IsEmpty() {
		return nil
	}
	v := cq.queue[cq.head]
	cq.head = (cq.head + 1) % cq.capacity
	return v
}

// 打印
func (cq *CircularQueue) Print() {
	if cq.head == cq.tail {
		fmt.Println("empty queue")
	} else {
		result := "head"
		for i := cq.head; i != cq.tail; i = (i + 1) % cq.capacity {
			result += fmt.Sprintf("<-%+v", cq.queue[i])
		}
		result += "<-tail"
		fmt.Println(result)
	}
}
