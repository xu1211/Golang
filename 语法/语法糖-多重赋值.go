package main
import "fmt"

func main() {
	var a int = 100
	var b int = 200
	// 变量的左值和右值按从左到右的顺序赋值。无需中间变量，即可变量交换
	b, a = a, b
	fmt.Println(a, b)
}