package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// *Vertex 指针类型的接收者
// 方法
/**
定义方式：有关联的接收者 (v *Vertex)
调用语法：v.Scale(10)（通过类型的实例调用）
与类型的关系：属于 *Vertex类型，是其行为的一部分
接口实现：可以实现接口
自动转换：Go 编译器会自动在值/指针间转换以方便调用（如 v.Scale(10)会被解释为 (&v).Scale(10)）
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//函数
/**
定义方式：无接收者，是独立的代码块
调用语法：ScaleFunc(&v, 10)（直接通过函数名调用）
与类型的关系：独立于任何类型，是一个通用工具
接口实现：不能直接用于实现接口
自动转换：参数类型必须严格匹配，需要显式传递指针（如 &v）
*/
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	fmt.Println(v)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

/**
带指针参数的函数必须接受一个指针
var v Vertex
ScaleFunc(v, 5)  // 编译错误！
ScaleFunc(&v, 5) // OK

接收者为指针的方法被调用的时候， 接收者既能是值 又能是指针
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
*/
