package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 10
	var b float64 = 3

	// if语句
	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 不小于 20\n")
	}

	// if语句 可以在条件之前执行一个简单语句
	if v := math.Mod(a, b); v < 10 {
		fmt.Printf("余数为：%v\n", v)
	}

	/*
	  switch
	*/
	// switch 多条件匹配
	switch a {
	case 1, 2, 3, 4:
		fmt.Printf("a 属于 1, 2, 3, 4\n")
	default:
		fmt.Printf("a 不属于 1, 2, 3, 4\n")
	}

	// switch 判断某个 interface 变量中实际存储的变量类型
	var x interface{}
	x = 1
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T\n", i)
	case int:
		fmt.Printf("x 是 int 型\n")
	case float64:
		fmt.Printf("x 是 float64 型\n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型\n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型\n")
	default:
		fmt.Printf("未知型\n")
	}

	// switch默认只会执行一个 case。如果case带有fallthrough，程序会继续执行下一条case，且不会去判断case表达式是否为true
	switch {
	case true:
		fmt.Println("1、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("2、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("3、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("4、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("5、默认 case")
	}

	/*
		select
		  select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测。
		  每个 case 必须是一个通信操作，发送/接收。
		  如果没有 case 可运行，它将阻塞，直到有 case 可运行
		  如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
	*/
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
