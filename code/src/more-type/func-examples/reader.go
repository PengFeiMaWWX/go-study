package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	//字符串读取器创建
	// strings.NewReader函数将字符串"Hello, Reader!"包装成一个strings.Reader对象，该对象实现了io.Reader接口
	//。这使得字符串可以像数据流一样被读取
	r := strings.NewReader("Hello, Reader!")

	//缓冲区初始化
	//创建一个大小为8字节的缓冲区。这个缓冲区用于每次从读取器中接收数据，8字节的大小意味着每次读取操作最多能获取8个字符
	b := make([]byte, 8)

	for {
		// r.Read(b)尝试将数据读入缓冲区b, 返回的n表示实际读取的字节数, err包含错误信息，正常时为nil，遇到流结束时为io.EOF
		n, err := r.Read(b)
		// 第一行显示读取字节数、错误信息和整个缓冲区内容
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		// 第二行使用b[:n]只显示实际读取的有效数据
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
