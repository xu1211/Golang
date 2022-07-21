package main

import (
	"fmt"
)

/*定义接口
    type 接口类型名 interface{
        方法名1( 参数列表1 ) 返回值列表1
        方法名2( 参数列表2 ) 返回值列表2
        …
    }
Go语言中实现接口是隐式的：实现接口中的 所有方法 , 就实现了这个接口
*/
type Phone interface {
	call()
	sms()
}

// 定义结构体MiPhone
type MiPhone struct {
	name string
	price float32
}
// 实现call()方法，就实现了Phone接口
func (miPhone MiPhone) call() {
	fmt.Println("I am xiaomi, I can call you!")
}

// 定义结构体IPhone
type IPhone struct {
	name string
	price float32
}
// 实现call()方法，实现Phone接口
func (iPhone *IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

// 结构体类型 MiPhone 和 *IPhone 都实现了 Phone接口，所以我们可以使用它们的实例作为 phoneCall() 的参数。
func phoneCall(p Phone) {
	fmt.Println(p)
	p.call()
}

func main() {
// 使用方法1：用 实例类型 作为 函数接口类型的参数
	mi := MiPhone{"小米01",1000}
	phoneCall(mi)
	i := IPhone{"苹果01",10000}
	// 由于 call() 只定义在 *IPhone（指针类型）上， 所以 只能用指针类型调用call()
	phoneCall(&i)

// 使用方法2：用接口类型 接收 实例类型
	var p Phone
	// p MiPhone 实现了call()
	p = new(MiPhone)
	p.call()
	// p *IPhone 实现了call()
	p = new(IPhone)
	p.call()


/*	// 由于 call() 只定义在 *IPhone（指针类型）上， 所以 IPhone（值类型）不能调用call()
	v := IPhone{}
	p = v 	//会报错， 写成 p = &v 才能调用call()
	p.call()
*/
}
