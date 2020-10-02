package main

import "fmt"

/*
定义结构体 Books
	Title ：标题
	Author ：作者
	Subject ：学科
	ID ：书籍ID

相对于包:
首字母大写相当于 public。
首字母小写相当于 private。

*/
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// 将结构体对象转换为 JSON 时，对象中的属性首字母必须是大写，才能正常转换为 JSON。
type Person struct {
	　　　Name string　　　　　　//Name字段首字母大写
	　　　age int               //age字段首字母小写
	}
	
func main() {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "Google", "Go 语言教程", 1})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "Google", subject: "Go 语言教程", book_id: 2})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "Google"})

	// 点号 . 操作符 访问结构体成员
	println()
	var Book1 Books
	Book1.title = "Go 语言"
	Book1.author = "Google"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 3
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	// 结构体作为函数参数
	println()
	printBook1(Book1)

	// 结构体指针
	println()
	printBook2(&Book1)
}

func printBook1(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook2(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
