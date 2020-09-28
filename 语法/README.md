# Golang
 Golang实践

- 行分隔符

在 Go 程序中，一行代表一个语句结束，一行想写多个语句则必须使用 `;` 区分

- [常量](./2variable.go)

- [变量](./3constant.go)
声明必须使用空格隔开,
```Go
// 单个变量声明
var v_name v_type
// 单个变量声明+赋值
var v_name v_type = value
// 多个变量声明+赋值
var v_name1, v_name2 v_type = value1, value2
// 根据值自行判定变量类型
var v_name = value
// 省略 var,只能在函数体中出现，:= 左侧的变量不能被声明过，会导致编译错误
v_name := value
```


如果没有初始化，则变量默认为零值。
- 数值类型（包括complex64/128）为 0
- bool 零值为 false
- 字符串为 ""（空字符串）


- [指针](./4指针.go)


- [条件控制](./5条件控制.go)
if...else语句, switch 语句,select 语句

- [循环](./6循环.go)
for语句, goto语句

- [函数](./7函数.go)
函数定义,调用,2种传值方式

- [函数使用方式1](./7函数1.go)
匿名函数 闭包直接使用函数内的变量
