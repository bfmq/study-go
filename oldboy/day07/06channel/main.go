package main

import "fmt"

func noBufChannel() {
	ch := make(chan int)
	go func() {
		x := <-ch
		fmt.Println(x)
	}()
	ch <- 10
	fmt.Println("10放进去了")
}

func bufChannel() {
	ch := make(chan int, 5)
	ch <- 10
	fmt.Println("10放进去了")
	x := <-ch
	fmt.Println(x)
	close(ch)
}

func main() {
	bufChannel()
}
