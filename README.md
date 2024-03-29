[toc]

# Golang
 Golang实践

## 基础语法
- [基础语法](./语法/README.md)
  - 常量,变量
  - 分支,循环
  - 数据结构,指针
  - 函数,方法 func
  - 接口 interface
  - 包 package
  - 异常 panic


### goroutine 并发处理
- [并发 `goroutine`](./语法/15并发.go)
> go语言中使用 协程(goroutine) 实现并发; 它是由官方实现的超级"线程池"。 由runtime管理
> 
> 内存占用: 每个 goroutine只占用4KB。 而线程要8MB; 
> 同一台机器 Java线程 最多只能建几千个，但是 Go协程 可以建数百万个。

- [通道 `channel`](./语法/16channel.go)

  功能: goroutine之间通讯

  channel有发送（send）、接收(receive）和关闭（close）三种操作。
- [通道channel- 缓存](./语法/17channel缓冲.go)

  `make(chan 元素类型, [缓冲大小])`
- [通道channel- 关闭,`for 循环接收`](./语法/18channel关闭与遍历.go)

  接收时判断channel是否关闭:`num, ok := <-ch`
- [通道channel- `range 自动关闭的循环接收`](./语法/18.1RangeChannel.go)

  channel关闭后,range会自动关闭接收
- [通道channel-`select 多通道接收`](./语法/19channelSelect.go)
    - [select案例-斐波那契数列](./语法/20channelSelectFibonacci.go)

  for-select 多通道循环接收

#### 上下文 `context 接口`
context是内置接口. 详细见下文: 内置接口-context

  简单案例: [用context停止子协程](./语法/context.go)
> 实际是 goroutine 的上下文，包含 goroutine 的运行状态、环境、现场等信息。
>
> 用来在 goroutine之间 传递上下文信息，包括：取消信号、超时时间、截止时间、k-v 等。


#### 同步,锁 `sync 包`
详细见下文: 内置包-sync包
#### cpu调度 `runtime 包`
详细见下文: 内置包-runtime包


### 内置接口
- [Stringer接口](./内置接口/14.3内置Stringer接口.go)
- [error接口](./内置接口/14.4内置error接口.go)
- [Readers接口](./内置接口/14.5内置Readers接口.go)
- [http接口：web服务实现](./内置接口/14.6内置http.go)\
  http://localhost:8080/
- [http接口：实现多个处理方法](./内置接口/14.7内置http2.go)\
  http://localhost:8081/string \
  http://localhost:8081/struct
- [image接口](./内置接口/14.8内置image.go)
#### context接口
Context是线程安全的; 
对于一条函数调用链 必须传递context上下文
- [接口方法; 创建context; 取消context; 传值](./内置接口/context.go)


## 内置包 package

- [runtime包](./内置包/runtime)
  运行时, cpu调度
  - [限制并发数](./内置包/runtime/MAXPROCS.go)
  - [让出cpu](./内置包/runtime/sched.go)
  - [终止当前goroutine](./内置包/runtime/exit.go)
  - runtime.GC()
- [sync包](./内置包/sync)
  - [once 单次操作](./内置包/sync/once.go)
  - [sync.Mutex 互斥锁, sync.RWMutex读写锁](./内置包/sync/mutex.go)
  - [sync.WaitGroup 同步信号量](./内置包/sync/wait_group.go)
- atomic包
原子操作
- os包
  系统的基本操作，如文件操作、目录操作、执行命令、信号与中断、进程、系统状态
- time包 
时间,日期,定时器
- math/big 
大数字精度计算
- regexp
正则表达式
- flag
  解析命令行参数

- reflect包 反射

- unsafe 包
不安全的指针, 放开了语言对指针的安全限制
## 性能分析 pprof
- [pprof-web页面](./pprof/http.go)


## 测试
测试用例文件名都为: `*_test.go`
### 单元测试
```bash
# 执行unit_test.go文件的所有用例
go test ./test/unit_test.go

# -v，可以让测试时显示详细的流程
go test -v  ./test/unit_test.go

# -run [name\正则]，只执行符合名称的用例
go test -v -run TestA ./test/unit_test.go
```
### 基准测试 benchmark
```bash
# -bench=[name\正则] 运行符合名称的基准测试用例
go test -v -bench=. ./test/benchmark_test.go

# -benchmem 参数以显示内存分配情况
go test -v -bench=Alloc -benchmem benchmark_test.go

```

# 核心概念
## 调度器

- M:N 模型：
  
  Go runtime 会负责 goroutine 的生老病死，从创建到销毁，都一手包办。
  
  Runtime 会在程序启动的时候，创建 `M 个线程`（CPU 执行调度的单位），

  之后创建的 `N 个 goroutine` 都会依附在这 M 个线程上执行。

- G-M-P 调度模型
  G: 代表一个 goroutine, 协程
  M: 取 machine 的首字母，内核线程
  P: 代表一个虚拟的 Processor, 处理器

m 需要获得 p 才能运行 g。
## 内存管理
Go 运行时的分配算法基于 tcmalloc，基本上没有碎片问题
- GC 垃圾回收