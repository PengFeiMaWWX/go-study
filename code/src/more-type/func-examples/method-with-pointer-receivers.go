package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Scale方法：状态修改器
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 状态读取器
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("缩放前：%+v，绝对值：%v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("缩放后：%+v，绝对值：%v\n", v, v.Abs())

}

/**
使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样会更加高效。
*/
