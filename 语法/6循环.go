package main

import "fmt"

func main() {
	// for 循环
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
		}
		fmt.Println("")
	}
	println()

	// for 循环 省略分号; 等同while
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	println()

	// for + goto
	for m := 1; m < 10; m++ {
		n := 1
	LOOP:
		if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}

	// defer延迟函数的执行直到上层函数返回
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")
	//延迟的函数调用被压入一个栈中。函数返回时会按照后进先出的顺序调用被延迟的函数调用。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
