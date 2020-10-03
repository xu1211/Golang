package main

import (
	"fmt"
)

// 定义接口
type Phone interface {
	call()
}

// 定义结构体
type MiPhone struct {
}
type IPhone struct {
}

func main() {
	var phone Phone

	phone = new(MiPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

func (miPhone MiPhone) call() {
	fmt.Println("I am xiaomi, I can call you!")
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
