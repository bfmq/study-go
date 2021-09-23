package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan chan bool = make(chan bool, 1)

func f() {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("========")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break FORLOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 通知该退出了
	exitChan <- true
	wg.Wait()
}
