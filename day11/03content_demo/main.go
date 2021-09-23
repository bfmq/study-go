package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1(ctx context.Context) {
	defer wg.Done()
	wg.Add(1)
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("f1111111111111111111111111")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("f222222222222222222222222222")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f1(ctx)
	time.Sleep(time.Second * 5)
	// 通知该退出了
	cancel()
	wg.Wait()
}
