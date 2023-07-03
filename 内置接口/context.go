package main

import (
	"context"
	"fmt"
	"time"
)

/*
context.Context 标准接口,有4个幂等的方法
*/
type Context interface {
	// 1. Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
	Deadline() (deadline time.Time, ok bool)

	// 2. Done — 返回一个 Channel，这个Channel读不出任何东西, 它会在当前工作完成或者上下文被取消后关闭；
	Done() <-chan struct{}

	// 3. Err — 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值；
	//	如果 context.Context 被取消，会返回 Canceled 错误；
	//	如果 context.Context 超时，会返回 DeadlineExceeded 错误；
	Err() error

	// 4. Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；
	Value(key interface{}) interface{}
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

// 5. 创建 带有相对截止时间的 Context;
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

func cancelChildContext() {
	// 创建一个父 Context
	parentCtx := context.Background()
	fmt.Printf("parent ctx: %v\n", parentCtx)
	// 创建一个子 Context
	childCtx, childCancel := context.WithCancel(parentCtx)

	// 在子 Context 中启动一个Goroutine协程
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("子 Context 取消")
				// 如果 Context 被取消，则退出 Goroutine
				return
			default:
				fmt.Println("子 Context的 Goroutine 运行中")
				time.Sleep(time.Second)
			}
		}
	}(childCtx)

	// 等待一段时间后取消子 Context
	time.Sleep(3 * time.Second)
	childCancel()

	// 等待一段时间后退出程序
	time.Sleep(time.Second)
	fmt.Printf("parent ctx 还在: %v\n", parentCtx)
}

func cancelParentContext() {
	// 创建一个父 Context
	parentCtx, parentCancel := context.WithCancel(context.Background())
	// 在父 Context 中创建一个子 Context
	childCtx, _ := context.WithCancel(parentCtx)

	// 在子 Context 中启动一个 Goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("子 Context 也被取消")
				// 如果 Context 被取消，则退出 Goroutine
				return
			default:
				fmt.Println("子 Context的 Goroutine 运行中")
				time.Sleep(time.Second)
			}
		}
	}(childCtx)

	// 等待一段时间后取消父 Context
	time.Sleep(3 * time.Second)
	parentCancel()

	// 等待一段时间后退出程序
	time.Sleep(time.Second)
}

func main() {
	fmt.Println("hello world")

	/*
	   创建context.Context 的6中创建方法
	*/
	fmt.Println("--------------------------------")
	//创建空的 Context
	backContext()
	fmt.Println("--------------------------------")
	//创建空的 Context，类似于 context.Background()，但更适用于暂时不确定使用哪种 Context 的情况
	todoContext()
	fmt.Println("--------------------------------")
	//创建 带有取消功能的 Context
	withCancel(context.Background())
	fmt.Println("--------------------------------")
	//创建 带有截止时间与取消功能的 Context
	WithDeadline(context.Background())
	fmt.Println("--------------------------------")
	//创建 带有相对截止时间的 Context
	withTimeout(context.Background())
	fmt.Println("--------------------------------")
	//创建 带有k-v的 Context
	withValue(context.Background())

	/* 取消context
	var cancel context.CancelFunc
	cancel() 方法的功能是: 关闭context.Done()；递归地取消它的所有子节点；从父节点从删除自己；
	*/
	fmt.Println("--------------------------------")
	//取消子context
	cancelChildContext()
	fmt.Println("--------------------------------")
	//取消父context, 其下所有子context也会被取消
	cancelParentContext()

	fmt.Println("end")
}
