package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}
type t struct {
	S string
}

func (t *t) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &t{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)

}
