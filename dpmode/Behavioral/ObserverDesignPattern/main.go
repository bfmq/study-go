package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Data int
}

// 观察者接口
type Observer interface {
	NotifyCallBack(event Event)
}

// 被观察者接口
type Subject interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

// 观察者
type ConcreteObserver struct {
	ID   int
	Time time.Time
}

// 回调事件
func (e *ConcreteObserver) NotifyCallBack(event Event) {
	fmt.Println(fmt.Sprintf("%dRecieved:%d after %v\n", e.ID, event.Data, time.Since(e.Time)))
}

// 被观察者
type ConcreteSubject struct {
	Observers *sync.Map
}

// 添加观察者
func (e *ConcreteSubject) AddListener(obs Observer) {
	e.Observers.Store(obs, struct{}{})
}

func (e *ConcreteSubject) RemoveListener(obs Observer) {
	e.Observers.Delete(obs)
}

// 通知各个观察者异步触发其回调函数
func (e *ConcreteSubject) Notify(event Event) {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		go key.(Observer).NotifyCallBack(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = j, i+j {
			out <- i
		}
	}()
	return out
}

func main() {
	n := ConcreteSubject{Observers: &sync.Map{}}
	obs1 := ConcreteObserver{
		ID:   1,
		Time: time.Now(),
	}
	obs2 := ConcreteObserver{
		ID:   2,
		Time: time.Now(),
	}
	n.AddListener(&obs1)
	n.AddListener(&obs2)
	for x := range Fib(10) {
		n.Notify(Event{Data: x})
	}
}
