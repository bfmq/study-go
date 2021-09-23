package main

import (
	"context"
	"fmt"
	"time"
)

func p(ctx context.Context, ch chan<- int) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- 1
			fmt.Println("生产者生产中...")
			time.Sleep(time.Second)
		}
	}
}

func c(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			i := <-ch
			fmt.Println("消费者消费中-------->", i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ch := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	go p(ctx, ch)
	go c(ctx, ch)
	time.Sleep(10 * time.Second)
}
