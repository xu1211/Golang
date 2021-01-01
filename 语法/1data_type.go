package main
import "fmt"
/*
基本类型：
    布尔型   bool
    数值类型
        有符号整数     int、int8、int16、int32、int64
        无符号整数     uint、uint8、uint16、uint32、uint64、uintptr
        浮点型   float32、float64
        复数    complex64、complex128
    字符串   string
    字符类型
        byte（=unit8）代表了 ASCII 码的一个字符。
        rune（=int32）代表一个 Unicode 字符
派生类型：
    指针类型（Pointer）
    数组类型（Array）
    切片类型（Slice）
    Map类型（Map）
    结构类型（Struct）
    管道类型（Channel）
    函数类型（Function）
    接口类型（Interface）
*/
func main() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)

    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}