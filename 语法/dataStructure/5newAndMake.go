/*
make和new关键字的区别及实现原理
new 和 make 是两个内置函数，主要用来创建并分配类型的内存

- new
func new(Type) *Type
    入参：接受一个（类型）参数
    作用：分配内存空间，空间值为类型的零值。
    出参：一个指向该类型内存地址的（指针）。

- make
func make(t Type, size ...IntegerType) Type
    入参：只能是 slice、map、channel 这三种引用类型
    作用：分配内存空间，初始化
    出参：chan、map、slice（类型本身）
*/
package main

import (
	"bytes"
	"fmt"
	"sync"
)

type Student struct {
	name   string
	age    int
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	fmt.Println("[paink] 内存空间为空不可使用 --------------------------------")
	paink1()

	fmt.Println("new 分配内存空间   --------------------------------")
	newTest()

	fmt.Println("[paink] slice, map, channel不能用new--------------------------------")
	paink2()

	fmt.Println("make --------------------------------")
	makeTest()

}

// paink1 : 对没有内存空间的变量无法赋值
func paink1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("对没有内存空间的变量直接赋值会paink, recover: [%s]  \n", r)
		}
	}()

	// 对已存在内存空间的变量赋值
	var temp int
	var v1 *int
	v1 = &temp
	fmt.Printf("int 内存地址: %p ,指针值: %#v ,指针指向的值: %#v \n", &v1, v1, *v1)
	*v1 = 11
	fmt.Printf("int赋值后,内存地址: %p ,指针值: %#v ,指针指向的值: %#v \n", &v1, v1, *v1)

	var student1 Student
	fmt.Println("Student 初始值: ", student1)
	student1.name = "s1" // struct 有内存空间,可赋值
	fmt.Println("Student 赋值: ", student1)

	// 对没有内存空间的变量赋值, 会paink
	var v2 *int // 指针的零值是nil
	fmt.Printf("*int 内存地址: %p ,指针值: %#v \n", &v2, v2)
	*v2 = 8 // *int 没有内存空间,会paink

	var s *Student
	fmt.Printf("*struct 内存地址: %p ,指针值: %#v \n", &s, s)
	s.name = "s2" // *struct 没有内存空间,会paink
}

// newTest : 用new分配 内存空间
func newTest() {
	// *int
	var v3 *int
	v3 = new(int) // 给*int分配空间
	fmt.Printf("new *int 内存地址: %p ,指针值: %#v ,指针指向的值: %#v \n", &v3, v3, *v3)
	*v3 = 98
	fmt.Printf("new *int赋值后,内存地址: %p ,指针值: %#v ,指针指向的值: %#v \n", &v3, v3, *v3)
	fmt.Println()

	// *struct
	var s *Student
	s = new(Student) //分配空间
	fmt.Printf("new *struct 内存地址: %p ,指针值: %#v ,指针指向的值: %#v \n", &s, s, *s)
	s.name = "dequan"
	s.age = 18
	fmt.Printf("new *struct赋值后, 内存地址: %p ,指针值: %#v ,指针指向的值: %#v \n", &s, s, *s)
	fmt.Println()
}

// paink2 : 引用类型初始化为nil，nil是不能直接赋值。并且不能用new分配内存
func paink2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("对没有内存空间的变量直接赋值会paink, recover: [%s]  \n", r)
		}
	}()

	// silce
	var a *[]int
	fmt.Printf("silce 内存地址: %p ,值为空: %#v \n", &a, a)
	av := new([]int)
	fmt.Printf("new silce 内存地址:  %p ,值还是为空: %#v \n", &av, av)
	(*av)[0] = 8 //panic: runtime error: index out of range

	// map
	var m map[string]string
	fmt.Printf("map 内存地址: %p, 值为空: %#v \n", &m, m)
	mv := new(map[string]string)
	fmt.Printf("new map 地址: %p, 值还是为空: %#v \n", &mv, mv)
	(*mv)["a"] = "a" //panic: assignment to entry in nil map

	// chan
	var c chan string
	fmt.Printf("chan 内存地址: %p ,值为空: %#v \n", &c, c)
	cv := new(chan string)
	fmt.Printf("new chan 内存地址: %p ,指针值: %#v ,指针指向的chan值还是为空: %#v \n", &cv, cv, *cv)
	*cv <- "string" //goroutine 1 [chan send (nil chan)]:
}

func makeTest() {

	av := make([]int, 5, 10)
	fmt.Printf("make silce 内存地址: %p ,值: %#v ,len:%#v ,cap:%#v \n", &av, av, len(av), cap(av))
	av[0] = 1 // 不报错,可使用

	mv := make(map[string]string)
	fmt.Printf("make map 内存地址: %p ,值: %#v \n", &mv, mv)
	mv["m"] = "m"

	chv := make(chan string)
	fmt.Printf("make chan 内存地址: %p ,值: %#v \n", &chv, chv)
	go func(message string) {
		chv <- message // 存消息
	}("Ping!")
	fmt.Println(<-chv) // 取消息 //"Ping!"
	close(chv)
}
