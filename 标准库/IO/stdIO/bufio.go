package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/**
	bufio 读取完整一行
	*/
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：\n")

	/**
	基础方法: 读取，直到遇到第一个界定符（delim）为止
	返回的 []byte 是指向 Reader 中的 buffer;  多次读取会覆盖
	*/
	slice, err := reader.ReadSlice('\n')
	fmt.Printf("%#v\n", slice)

	// 等于 ReadSlice + copy值; 多次读取不会覆盖
	buf, err := reader.ReadBytes('\n')
	fmt.Printf("%#v\n", buf)

	// 等于 ReadBytes('\n')
	line, _, err := reader.ReadLine()
	if err == io.EOF {
		fmt.Printf("%#v\n", line)
	}
	fmt.Printf("%#v\n", line)

	// 等于 ReadBytes + 转string
	text, err := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)

}
