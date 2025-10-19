package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat64 float64

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

func main() {

	var a Abser
	f := MyFloat64(-math.Sqrt2)
	a = f // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())

	v := Vertex{3, 4}
	a = &v // a *Vertex 实现了 Abser
	fmt.Println(a.Abs())

}

/**

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f *MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

这两个方法有什么区别

​修改原始值的能力​

​值接收者 (f MyFloat)​​：方法内部操作的是接收者值的副本。对f的任何修改（例如 f = 10）都只会影响这个副本，而不会改变原始的MyFloat变量
。

​指针接收者 (f *MyFloat)​​：方法内部操作的是指向原始值的指针。因此，通过指针解引用（例如 *f = 10）可以直接修改原始的MyFloat变量
。不过，在你提供的Abs方法中，并没有对f进行赋值修改，所以这一点在此例中体现不明显。

​性能考量​

对于MyFloat这样的基本类型（占用内存小，通常为8字节），使用值接收者复制副本的开销非常小，甚至可以忽略。此时，两种方式在性能上几乎没有差异。

但如果接收者是一个大型结构体​（包含很多字段），使用值接收者会在每次方法调用时复制整个结构体，产生明显的性能开销。在这种情况下，使用指针接收者（只复制指针，通常为8字节）会更高效​
。

​接口实现的规则​

这是两者一个关键且容易混淆的区别，涉及到Go语言的方法集规则
：

​值类型（MyFloat）​​ 的方法集只包含值接收者声明的方法。

​指针类型（*MyFloat）​​ 的方法集包含值接收者和指针接收者声明的所有方法。

这意味着：

如果Abs方法是某个接口（例如Abser）必须实现的方法，那么：

使用值接收者时，MyFloat类型和*MyFloat类型都实现了该接口。

使用指针接收者时，只有*MyFloat类型实现了该接口。如果你尝试将一个MyFloat值（而非指针）赋值给Abser接口变量，编译器会报错
。
*/
