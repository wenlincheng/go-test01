package main

import (
	"fmt"
	"math"
)

type Girl struct {
	Height, Weight float64
}

// 函数形式
func Abs(g Girl) float64 {
	return math.Sqrt(g.Height*g.Height + g.Weight*g.Weight)
}

func main() {
	g := Girl{3, 4}
	fmt.Println(Abs(g))
}
