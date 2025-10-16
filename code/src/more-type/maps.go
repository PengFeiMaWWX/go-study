package main

import "fmt"

var m map[string]Vertex

type Vertex struct {
	Lat, Lang float64
}

func main() {
	m = make(map[string]Vertex)
	fmt.Println(m["Bell Labs"])
	m["A B"] = Vertex{3, 4}
	fmt.Println(m["A B"])
}
