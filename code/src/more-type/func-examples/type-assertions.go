package main

import "fmt"

func main() {

	/**
	以把 interface{}想象成一个万能容器或者一种通用类型。正因为它是“空”的（没有规定任何必须实现的方法），
	所以 Go 语言中的所有类型，无论是基本的 int, string，还是复杂的结构体或切片，都天然地符合它的要求，可以被放进去
	*/
	var i interface{} = "hello"
	// 断言 // 不安全的方式：如果断言失败，程序会触发panic（运行时错误）
	s := i.(string)
	fmt.Println(s)
	// 安全的方式：使用两个返回值，ok为true表示断言成功
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
