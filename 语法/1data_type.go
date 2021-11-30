package main

import (
	"fmt"
	"unsafe"
)

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
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Println("int: ", unsafe.Sizeof(i1))
	fmt.Println("int: ", unsafe.Sizeof(i2))
	fmt.Println("int: ", unsafe.Sizeof(i3))
	fmt.Println("int: ", unsafe.Sizeof(i4))
	fmt.Println("int: ", unsafe.Sizeof(i5))

	var ui1 uint = 1
	var ui2 uint8 = 2
	var ui3 uint16 = 3
	var ui4 uint32 = 4
	var ui5 uint64 = 5
	fmt.Println("uint: ", unsafe.Sizeof(ui1))
	fmt.Println("uint8: ", unsafe.Sizeof(ui2))
	fmt.Println("uint16: ", unsafe.Sizeof(ui3))
	fmt.Println("uint32: ", unsafe.Sizeof(ui4))
	fmt.Println("uint64: ", unsafe.Sizeof(ui5))

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)

	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
