package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup

// var lock sync.Mutex

func add() {
	defer wg.Done()
	// lock.Lock()
	// defer lock.Unlock()
	atomic.AddInt64(&x, 1)
	// x++
}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
