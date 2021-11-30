package main

import "fmt"

/*
定义结构体 Books
	Title ：标题
	Author ：作者
	Subject ：学科
	ID ：书籍ID

相对于包:
	首字母大写 相当于 public。
	首字母小写 相当于 private。
*/
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 实例化 结构体
	pythonbook1 := Books{
		"python 入门",
		"py",
		"编程",
		1}
	fmt.Println("结构体实例化方式1：", pythonbook1)

	// 也可以使用 key => value 格式
	pythonbook2 := Books{
		title:   "python 放弃",
		author:  "py",
		subject: "编程",
		book_id: 2}
	fmt.Println("结构体实例化方式2：", pythonbook2)

	// 省略的字段将被初始化为零值
	var gobook Books
	fmt.Println("结构体实例化方式3：", gobook)

	// 点号 . 操作符 访问结构体成员
	gobook.title = "Go 语言"
	gobook.author = "Google"
	fmt.Println("访问结构体成员：", gobook.title)
	println()

	// 指针类型的结构体： new 实例化
	cbook1 := new(Books)
	cbook1.title = "C 语言 上"
	fmt.Println("指针类型的结构体：", cbook1)
	fmt.Println("结构体指针使用.会被自动解引用：", cbook1.title)
	fmt.Println("指针类型的结构体：", *cbook1)

	// 指针类型的结构体：对结构体进行 &取地址操作
	cbook2 := &Books{}
	cbook2.title = "C 语言 下"
	fmt.Println("指针类型的结构体：", cbook2)
	fmt.Println("指针类型的结构体：", *cbook2)

	// 结构体 作为函数参数
	println()
	printBook1(gobook)

	// 结构体指针 作为函数参数
	println()
	printBook2(cbook1)
}

func printBook1(book Books) {
	fmt.Println("Book title : ", book.title)
	fmt.Println("Book author : ", book.author)
	fmt.Println("Book subject : ", book.subject)
	fmt.Println("Book book_id : ", book.book_id)
}

// 结构体指针使用 点号 . 操作符 指针会被自动解引用
func printBook2(book *Books) {
	fmt.Println("Book title : ", book.title)
	fmt.Println("Book author : ", book.author)
	fmt.Println("Book subject : ", book.subject)
	fmt.Println("Book book_id : ", book.book_id)
}
