package main

import (
	"context"
	"fmt"
	"time"
)

/*
context.Context 标准接口,有4个方法
*/
type Context interface {
	// 1. Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
	Deadline() (deadline time.Time, ok bool)

	// 2. Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 Done 方法会返回同一个 Channel；
	Done() <-chan struct{}

	// 3. Err — 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值；
	//	如果 context.Context 被取消，会返回 Canceled 错误；
	//	如果 context.Context 超时，会返回 DeadlineExceeded 错误；
	Err() error

	// 4. Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；
	Value(key interface{}) interface{}
}

/*
context.Context 的5中创建方法
*/
func main() {
	fmt.Println("hello world")

	backContext()
	todoContext()
	withCancel(context.Background())
	WithDeadline(context.Background())
	withTimeout(context.Background())
	withValue(context.Background())

	fmt.Println("end")
}

// 1. 创建空的 Context
func backContext() {
	ctx := context.Background()
	fmt.Println("Background: ", ctx)
}

// 2. 创建空的 Context，类似于 context.Background()，但更适用于暂时不确定使用哪种 Context 的情况
func todoContext() {
	ctx := context.TODO()
	fmt.Println("TODO: ", ctx)
}

// 3. 创建 带有取消功能的 Context
func withCancel(parentCtx context.Context) {
	// 在父 Context 中创建一个带有 Cancel 的子Context
	ctx, cancel := context.WithCancel(parentCtx)
	fmt.Printf("WithCancel ctx: %v, cancelFunc: %v \n", ctx, cancel)

	// 在子协程中 取消 Context
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	// 监听 ctx.Done()
	select {
	case <-ctx.Done():
		// 打印 ctx.Err()
		fmt.Printf("withCancel err: %v\n", ctx.Err())
		if ctx.Err() == context.Canceled {
			fmt.Println("withCancel 上下文已取消")
		}
	}
}

// 4. 创建 带有截止时间与取消功能的 Context
func WithDeadline(parentCtx context.Context) {
	// 在父 Context 中创建一个带有 Timeout 的子Context
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(parentCtx, deadline)

	defer cancel() // 确保在函数结束时取消 Context

	// 监听 ctx.Done()
	select {
	case <-ctx.Done():
		// 打印 ctx.Err()
		fmt.Printf("WithDeadline err: %v\n", ctx.Err())
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("WithDeadline 上下文已超时")
		} else if ctx.Err() == context.Canceled {
			fmt.Println("WithDeadline 上下文已取消")
		}
	}
}

// 5. 创建 带有截止时间的 Context;
// 实际上是 context.WithDeadline() 方法的一个封装, 区别在于 超时时间是相对时间: time.Now().Add(timeout)
func withTimeout(parentCtx context.Context) {
	// 在父 Context 中创建一个带有 Timeout 的子Context
	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)

	defer cancel() // 确保在函数结束时取消 Context

	select {
	case <-ctx.Done():
		// 打印 ctx.Err()
		fmt.Printf("withTimeout err: %v\n", ctx.Err())
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("withTimeout 上下文已超时")
		} else if ctx.Err() == context.Canceled {
			fmt.Println("withTimeout 上下文已取消")
		}
	}
}

// 6. 创建 带有k-v的 Context
func withValue(parentCtx context.Context) {
	// 在父 Context 中创建一个带有 Value 的子 Context
	ctx := context.WithValue(parentCtx, "key", "value")

	// 在 子Context 中获取 Value
	value := ctx.Value("key")
	fmt.Println("withValue key: ", value)
}
