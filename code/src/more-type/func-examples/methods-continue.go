package main

import (
	"fmt"
	"math"
)

type MyFloat float64

/*
*
接收者的类型定义和方法声明必须在同一包内
*/
func (f MyFloat) Abs() float64 {

	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func main() {
	//​math.Sqrt2​：这是Go语言math包中预定义的一个常数，表示数学常数√2（约等于1.4142135623730951）
	// ​-math.Sqrt2​：对math.Sqrt2取负值，得到约等于-1.4142135623730951
	// ​MyFloat(...)​：这是一个类型转换，将-math.Sqrt2（默认是float64类型）转换为自定义类型MyFloat
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
