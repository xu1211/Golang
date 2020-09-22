# Golang
 Golang实践

- 行分隔符

在 Go 程序中，一行代表一个语句结束，一行想写多个语句则必须使用 `;` 区分

- 变量的声明必须使用空格隔开,
```Go
// 单个变量声明
var v_name v_type
// 单个变量声明+赋值
var v_name v_type = value
// 多个变量声明+赋值
var v_name1, v_name2 v_type = value1, value2
// 根据值自行判定变量类型
var v_name = value
// 省略 var,:= 左侧的变量不能被声明过的，会导致编译错误
v_name := value
```


如果没有初始化，则变量默认为零值。
- 数值类型（包括complex64/128）为 0
- bool 零值为 false
- 字符串为 ""（空字符串）