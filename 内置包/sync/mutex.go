package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 并发问题
	concurrencyProblem()
	fmt.Println()

	// 互斥锁
	mutex()
	fmt.Println()

	// 读写锁
	rwmutex()
	fmt.Println()

}

func concurrencyProblem() {
	var a = 0
	for i := 0; i < 100; i++ {
		// 协程非顺序执行
		go func(idx int) {
			a += 1
			fmt.Println(a)
		}(i)
	}
	time.Sleep(time.Second)
}

/*
一个互斥锁只能同时被一个 goroutine 锁定，
其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）
*/
func mutex() {
	// 互斥锁
	var lock sync.Mutex

	var a = 0
	for i := 0; i < 10; i++ {
		go func(idx int) {
			fmt.Printf("goroutine %d 等待解锁\n", idx)
			lock.Lock() // 加锁, goroutine获取锁才能进入临界区
			fmt.Printf("goroutine %d 抢到锁\n", idx)
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
			fmt.Printf("goroutine %d 解锁\n", idx)

			defer lock.Unlock() // 解锁
		}(i)
	}
	// 等待 1s 结束主程序
	// 确保所有协程执行完
	time.Sleep(time.Second)
}

/*
可以有任意多个 gorouinte 获得 读锁
只能有一个 goroutine 能够获得 写锁
读锁和写锁互斥, 同时只能存在一个
*/
var rwlock sync.RWMutex
var count int

func rwmutex() {
	// 读写锁
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
func read(n int, ch chan struct{}) {
	rwlock.RLock() // 加读锁
	fmt.Printf("goroutine %d 获取到读锁...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rwlock.RUnlock() // 解读锁
	ch <- struct{}{}
}
func write(n int, ch chan struct{}) {
	rwlock.Lock() // 加写锁
	fmt.Printf("goroutine %d 获取到写锁...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rwlock.Unlock() // 解写锁
	ch <- struct{}{}
}
