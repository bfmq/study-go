package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func f() {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 5; i++ {
// 		r1 := rand.Int()
// 		r2 := rand.Intn(10)
// 		fmt.Println(r1)
// 		fmt.Println(r2)
// 	}
// }

var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(30)))
	fmt.Println(i)
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go f1(i)
	}
	wg.Wait()
}
