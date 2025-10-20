package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

/**
 Person结构体定义的 String() string方法，签名与 Stringer接口的要求完全一致。
在 Go 语言中，接口的实现是隐式的
。这意味着只要一个类型拥有了接口所声明的全部方法，它就自动实现了该接口，
无需像其他语言一样使用 implements关键字显式声明。因此，你的 Person结构体已经实现了 fmt.Stringer接口。

​自动调用​：fmt包中的打印函数（如 fmt.Println）设计得非常智能。
在打印一个值时，它会检查该值的类型是否实现了 Stringer接口。如果实现了，
fmt包就会自动调用其 String()方法来获取该值的字符串表示，而不是使用默认的输出格式
*/
