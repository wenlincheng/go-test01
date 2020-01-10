package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell"])

	m2 := make(map[string]int)

	m2["a"] = 1
	m2["b"] = 2
	m2["c"] = 3

	fmt.Println(m2)
}
