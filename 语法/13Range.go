package main

import (
	"fmt"
	"os"
)

func main() {
	nums := []int{2, 3, 4}

	// 1.Range 简单循环
	length := 0
	for range nums {
		length++
		fmt.Println(length)
	}
	println()

	// 2.Range 循环array键值对
	println("Range 循环array键值对:")
	for i, num := range nums {
		fmt.Printf("array键：%d, 值：%d\n", i, num)
	}
	println()

	// range求和
	println("range 求和:")
	sum := 0
	// 空白符"_"省略元素的键
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	println()

	// range 循环map键值。
	println("Range 循环map键值对:")
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("map键：%s, 值：%s\n", k, v)
	}
	println()

	// range 循环枚举Unicode字符串
	println("Range 循环枚举键值对:")
	for k, v := range "go" {
		fmt.Printf("枚举键：%d, 值：%d\n", k, v)
	}
	println()

	println("range 获取参数列表:")
	fmt.Println(len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
	println()
}
