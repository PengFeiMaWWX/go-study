package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

/**
接受一个值作为参数的函数必须接受一个指定类型的值：

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！

而以值为接收者的方法被调用时，接收者既能为值又能为指针
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK

var v int = 42    // 声明一个整型变量 v
ptr := &v         // 使用 & 获取 v 的地址，并赋值给指针变量 ptr
fmt.Println(ptr)  // 输出类似 0xc0000140a8 的地址值
fmt.Println(*ptr) // 使用 * 解引用指针，输出 42（v 的值）
*/
