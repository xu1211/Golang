/*
make和new关键字的区别及实现原理
new 和 make 是两个内置函数，主要用来创建并分配类型的内存

- new
func new(Type) *Type
    入参：接受一个（类型）参数
    作用：分配内存空间，空间置为类型的零值。
    出参：一个指向该类型内存地址的（指针）。

- make
func make(t Type, size ...IntegerType) Type
    入参：只能是 slice、map 和 channel类型
    作用：分配内存空间，初始化
    出参：chan、map、slice（类型本身）
*/
package main
import "fmt"

func main() {
    // new分配 系统类型空间
    var sum *int
    sum = new(int) //分配空间
    *sum = 98
    fmt.Println(*sum)

    // new分配 自定义类型空间
    type Student struct {
        name string
        age int
    }
    var s *Student
    s = new(Student) //分配空间
    s.name ="dequan"
    fmt.Println(s)
}