package _9_queue

import "fmt"

type ArrayQueue struct {
	queue    []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		queue:    make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// 入队
func (aq *ArrayQueue) EnQueue(v interface{}) bool {
	if aq.tail == aq.capacity {
		return false
	}
	aq.queue[aq.tail] = v
	aq.tail++
	return true
}

// 出队
func (aq *ArrayQueue) DeQueue() interface{} {
	if aq.head == aq.tail {
		return nil
	}
	v := aq.queue[aq.head]
	aq.head++
	return v
}

// 打印
func (aq *ArrayQueue) Print() {
	if aq.head == aq.tail {
		fmt.Println("empty queue")
	} else {
		result := "head"
		for i := aq.head; i < aq.tail-1; i++ {
			result += fmt.Sprintf("<-%+v", aq.queue[i])
		}
		result += "<-tail"
		fmt.Println(result)
	}
}
