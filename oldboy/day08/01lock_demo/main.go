package main

import (
	"fmt"
	"sync"
)

var (
	x    = 0
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
