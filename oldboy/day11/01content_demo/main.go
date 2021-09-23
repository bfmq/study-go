package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify = false

func f() {
	defer wg.Done()
	for {
		fmt.Println("========")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 通知该退出了
	notify = true
	wg.Wait()
}
