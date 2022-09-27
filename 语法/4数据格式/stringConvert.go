package main

import "fmt"

func main() {
	str := "abcdefghijklmnopq"

	//string 转[]byte
	b := []byte(str)
	fmt.Printf("%#v\n", b)

	//[]byte转string
	str = string(b)
	fmt.Printf("%#v\n", str)

	//string 转 rune
	r := []rune(str)
	fmt.Printf("%#v\n", r)

	//rune 转 string
	str = string(r)
	fmt.Printf("%#v\n", str)
}
