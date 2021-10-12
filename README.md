# Golang
 Golang实践

- [基础语法](./语法/README.md)

## 内置接口
- [Stringer接口](./内置接口/14.3内置Stringer接口.go)
- [error接口](./内置接口/14.4内置error接口.go)
- [Readers接口](./内置接口/14.5内置Readers接口.go)
- [http接口：web服务实现](./内置接口/14.6内置http.go)\
  http://localhost:8080/
- [http接口：实现多个处理方法](./内置接口/14.7内置http2.go)\
  http://localhost:8081/string \
  http://localhost:8081/struct
- [image接口](./内置接口/14.8内置image.go)

## 并发处理
- [并发](./语法/15并发.go)
>线程的内存占用，每个 Go 线程竟然只占用4KB。同一台机器，Java 线程最多只能建几千个，但是 Go 线程可以建数百万个。
- [通道channel](./语法/16channel.go)
  多线程通讯
- [通道channel-缓存](./语法/17channel缓冲.go)
- [通道channel-关闭](./语法/18channel关闭与遍历.go)
- [通道channel-select语法](./语法/19channelSelect.go)
    - [select案例-斐波那契数列](./语法/20channelSelectFibonacci.go)



