package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  //创建一个Vertex类型的结构体
	v2 = Vertex{X: 1}  // Y 被隐式的赋值为0
	v3 = Vertex{}      //X:0 , Y:0
	p  = &Vertex{1, 2} // 创建一个Vertex类型的结构体（指针）
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
