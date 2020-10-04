package main

import (
	"fmt"
)

// 定义接口
type Phone interface {
	call()
}

// 定义结构体MiPhone 与 结构体方法
type MiPhone struct {
}

func (miPhone MiPhone) call() {
	fmt.Println("I am xiaomi, I can call you!")
}

// 定义结构体IPhone 与 结构体方法
type IPhone struct {
}

func (iPhone *IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var p Phone

	// p MiPhone 实现了call()
	p = new(MiPhone)
	p.call()

	// p *IPhone 实现了call()
	p = new(IPhone)
	p.call()

	// 由于 call() 只定义在 *IPhone（指针类型）上， 所以 IPhone（值类型）所以没有实现call()
	v := IPhone{}
	p = v // 写成 p = &v 才能实现call()
	p.call()
}
