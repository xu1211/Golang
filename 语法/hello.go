/*
   声明包
   package main表示一个可独立执行的程序
*/
package main

/* 引入 fmt 包 */
import "fmt"

/* 函数 */
func main() {
	fmt.Println("Hello, World!")
}

/*
    直接运行程序
    go run ./hello.go

    编译成二进制文件
    go build ./hello.go

    运行二进制文件
    hello.exe
*/