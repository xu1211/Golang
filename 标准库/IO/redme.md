
io 包最重要的是两个接口：Reader 和 Writer 接口

只要满足这两个接口，它就可以使用 IO 包的功能

os.Stdin/Stdout 它们分别实现了 io.Reader/io.Writer 接口
os.File 同时实现了这两个接口

# stdin输入
1. 只获取值,去除空格,换行符
   1. fmt.Scan	func Scan(a ...interface{}) (n int, err error) 	按空白符/换行符分隔stdin获取值, 传入到入参中
   2. fmt.Scanf	func Scanf(format string, a ...interface{}) (n int, err error)	只读取format格式的输入,其他格式都报错, 遇到换行时停止扫描, 不能换行
   3. fmt.Scanln	func Scanln(a ...interface{}) (n int, err error)	遇到换行时停止扫描, 不能换行


2. 全部获取, 包括空格
   1. bufio.NewReader


# stdout输出