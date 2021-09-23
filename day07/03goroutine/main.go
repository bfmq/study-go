package main

import (
	"fmt"
	"time"
)

// func hello(i int) {
// 	fmt.Println("hello", i)
// }

func main() {
	for i := 0; i < 1000; i++ {
		// go hello(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main")
	time.Sleep(time.Second)
}
