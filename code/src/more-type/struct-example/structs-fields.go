package main

import "fmt"

/*
*
结构体字段
结构体字段可通过点号 . 来访问。
*/
type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 5
	fmt.Println(v.X)
}
