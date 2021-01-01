package main

import (
	"fmt"
	"os"
	"io"
)

/*
隐式接口
	方式：实现接口中的所有方法,就实现了这个接口。 没有显式声明的必要；所以也就没有关键字“implements“。
	目的：解藕。 使 实现接口的包 和 定义接口的包 互不依赖。不需要再去画祖宗八代的继承关系图
		无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
但是隐式接口会带来冲突问题.
*/

//包 io 定义了 `Reader` 和 `Writer`
/*type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}
*/
func main() {
	// io.Writer 接口变量
	var w io.Writer

	// os.Stdout 实现了 io.Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
