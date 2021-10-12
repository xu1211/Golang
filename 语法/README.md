[@toc]

# Golang
 Golang学习实践


## 基础

- [HelloWord](./hello.go)

>行分隔符:
在 Go 程序中，一行代表一个语句结束，一行想写多个语句则必须使用 `;` 区分

>go 不支持隐式转换类型

- [**数据类型**](./1data_type.go)

- [变量 `var`](./2variable.go)
```go
var {identifier} {type}
```

零值
- [常量 `const`](./3constant.go)
```go
const {identifier} [{type}] = {value}
```
特殊常量 iota

- [语法糖-多重赋值](./语法糖-多重赋值.go)\
    多重赋值：无需中间变量，就能变量交换

## 逻辑控制
- [分支](./5条件控制.go)\
`if...else`语句, `switch` 语句,`select` 语句

- [循环](./6循环.go)\
`for`语句, `goto`语句
    - [`Range` 迭代](./7Range.go)\
    结合for使用，遍历数组集合

## 数据容器
- [指针](./4指针.go)

- 容器：定义/实例化/赋值/访问/传参
    - [数组 `[len]value_type`](./2容器/1Array.go)\
        长度固定，类型固定
    - [slice切片(动态数组) `[]value_type`](./2容器/2Slice.go)\
        长度不固定，类型固定；相对于数组 slice使用的更多，\
        append() \
        copy()
    - [map `map[key_type]value_type`](./2容器/4map.go)\
        无序的键值对的集合,其他的语言中称为 哈希 / 字典
    - [结构体 `struct`](./2容器/3.1struct.go)\
        结构体，指针类型的结构体    
        - [结构体对象转换为JSON](./2容器/3.2struct.go)
- [new() 与 make()](./2容器/5newAndMake.go)


## 函数,方法,接口
- [`func`：函数](./3函数/0函数.go)\
    函数声明,调用，
    值传递/引用传递，
    可变参数，
    多值返回，
    命名返回值
    - [函数变量](./3函数/1函数变量.go)
        **函数作为值 赋给变量**
    - [匿名函数](./3函数/2函数闭包.go)
        **闭包**
    - [函数回调](./3函数/4回调.go)
    - [函数递归](./3函数/5递归.go) 自己调用自己
    - [延迟函数`defer`](./3函数/6延迟函数.go)
      类似命令压栈，最后执行
        - [数组,切片作为参数](./3函数/7arrayToFunc.go)

- [`func`：方法](./3函数/3函数与方法.go)
    - 函数: `func function_name( [parameter list] ) [return_types] {函数体}`
    - 方法: `func (var receiver) method_name( [parameter list] ) [return_types] {方法体}`
>Go 没有面向对象，没有类。仍然可以在结构体类型上定义方法

- [变量作用域](./3函数/作用域.go)

- [接口 `interface`](./3函数/8接口.go)\
go里面没有类和继承的概念。
接口类型是由一组`方法定义的集合`。
接口类型的值可以存放实现这些方法的值    
    - [隐式实现接口](./3函数/9隐式接口.go)\
Go语言中**实现接口是隐式的**：实现接口中的 所有方法 , 就实现了这个接口

- [宕机 `panic`](./3函数/10程序宕机.go)
- [宕机恢复 `recover`](./3函数/11宕机恢复.go)