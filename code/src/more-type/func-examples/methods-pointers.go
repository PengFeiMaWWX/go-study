package main

/**
使用方法
*/
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
*
指针接收者的方法可以修改接收者指向的值（如这里的 Scale 所示）。
由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
值接收者方法​（如Abs()方法）：

有明确的返回类型float64

通过return语句返回计算结果

不修改原始对象的状态

​指针接收者方法​（如Scale(f float64)方法）：

确实没有声明返回类型

这是一个"无返回值"的方法，其目的是修改接收者对象的状态而非返回数据


*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())

}
