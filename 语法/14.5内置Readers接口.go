package main

/*
io 包指定了 io.Reader 接口， 它表示从数据流结尾读取。
Go 标准库包含了这个接口的许多实现， 包括文件、网络连接、压缩、加密等
io.Reader 接口有一个 Read 方法：func (T) Read(b []byte) (n int, err error)
Read 用数据填充指定的字节 slice，并且返回填充的字节数和错误信息。 在遇到数据流结尾时，返回 io.EOF 错误。
*/
import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	// Read读取
	for {
		n, err := s.Read(b)
		fmt.Printf("读取字符数n = %v 错误信息err = %v 数据b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	s1 := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s1}
	io.Copy(os.Stdout, &r)
}

// io.Reader 包裹另一个 `io.Reader`，然后通过某种形式修改数据流。
// 例：rot13解密
type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 'A' && b[i] < 'N':
			b[i] += 13
		case b[i] >= 'N' && b[i] <= 'Z':
			b[i] -= 13
		case b[i] >= 'a' && b[i] < 'n':
			b[i] += 13
		case b[i] >= 'n' && b[i] <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}
