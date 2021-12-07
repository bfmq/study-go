package main

import (
	"fmt"
	"time"
)

func f1(ch chan int, x, y, z int) {
	ch <- x + y + z
}

func main() {
	result := make([]int, 0, 2)
	ch := make(chan int, 3)
	go f1(ch, 1, 2, 3)
	go f1(ch, 4, 5, 6)

	for {
		select {
		case r := <-ch:
			result = append(result, r)
			if len(result) == 2 {
				close(ch)
				return
			}
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("==============")
		}
	}
}
