package main

import (
	"fmt"
	"time"
)

/*
`error` 类型是一个内建接口：使用 error 值来表示错误状态

type error interface {
    Error() string
}
*/

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at \n 时间：%v,\n 事件：%s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"未知错误",
	}
}

func main() {
	// 函数会返回一个 error 值，error 为 nil 时表示成功；非 nil 的 error 表示错误。
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
