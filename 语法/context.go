package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
例: 用context控制子协程停止
*/
var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		// 执行协程
		fmt.Println("worker")
		time.Sleep(time.Second)

		// 接收外部取消信号
		select {
		case <-ctx.Done(): // 协程中获取context的done channel
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	// 1. 创建context
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)

	// context传给其他goroutine
	go worker(ctx)

	// 通知子goroutine结束
	time.Sleep(time.Second * 3)
	cancel()

	// 等待子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
