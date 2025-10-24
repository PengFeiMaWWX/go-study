package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// 为MyReader 添加 Read方法
func (r MyReader) Read(b []byte) (int, error) {
	// 逐个遍历切片 b中的每个元素
	for i := range b {
		// 为缓冲区每个字节设置一个字母A
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
