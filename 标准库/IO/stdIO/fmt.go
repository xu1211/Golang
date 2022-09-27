package main

import (
	"fmt"
)

func main() {

	/**
	fmt库
	扫描stdin只获取值,去除空格,换行符
		fmt.Scan	func Scan(a ...interface{}) (n int, err error) 	按空格符/换行符分隔stdin获取值, 传入到入参中
		fmt.Scanf	func Scanf(format string, a ...interface{}) (n int, err error)	只读取format格式的输入,其他格式都报错, 遇到换行时停止扫描, 不能换行
		fmt.Scanln	func Scanln(a ...interface{}) (n int, err error)	遇到换行时停止扫描, 不能换行
	*/
	var (
		name    string
		age     int
		married bool
		ag      int
		err     error
	)

	ag, err = fmt.Scan(&name, &age, &married)
	fmt.Print(err)  // 这里如果出错，则显示出错的原因
	fmt.Println(ag) // 这个返回参数取到的是成功了几个
	fmt.Printf("Scan扫描结果1 name:%s age:%d married:%t \n", name, age, married)
	fmt.Scanln() //吸收换行符

	ag, err = fmt.Scan(&name, &age, &married)
	fmt.Print(err)  // 这里如果出错，则显示出错的原因
	fmt.Println(ag) // 这个返回参数取到的是成功了几个
	fmt.Printf("Scan扫描结果2 name:%s age:%d married:%t \n", name, age, married)
	fmt.Scanln() //吸收换行符

	ag, err = fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Print(err)
	fmt.Println(ag)
	fmt.Printf("Scanf扫描结果 name:%s age:%d married:%t \n", name, age, married)
	fmt.Scanln() //吸收换行符

	ag, err = fmt.Scanln(&name, &age, &married)
	fmt.Print(err)
	fmt.Println(ag)
	fmt.Printf("Scanln扫描第一行 name:%s age:%d married:%t \n", name, age, married)

	ag, err = fmt.Scanln(&name, &age, &married)
	fmt.Print(err)
	fmt.Println(ag)
	fmt.Printf("Scanln扫描第二行 name:%s age:%d married:%t \n", name, age, married)
}
