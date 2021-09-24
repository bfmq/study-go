package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []int
	cond  *sync.Cond
}

func (q *Queue) Enqueue(i, num int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, i)
	fmt.Printf("生产者%d生成了消息\n", num)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.queue) == 0 {
		q.cond.Wait()
	}
	i := q.queue[0]
	q.queue = q.queue[1:]
	return i
}

func main() {
	q := &Queue{
		queue: []int{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	for num := 1; num < 4; num++ {
		go func(num int) {
			for {
				q.Enqueue(1, num)
				time.Sleep(2 * time.Second)
			}
		}(num)
	}

	for num := 1; num < 4; num++ {
		go func(num int) {
			for {
				i := q.Dequeue()
				fmt.Printf("消费者%d消费了消息%d\n", num, i)
			}
		}(num)
	}

	select {}
}
