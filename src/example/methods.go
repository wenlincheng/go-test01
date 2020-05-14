package main

import (
	"fmt"
	"math"
)

type Phone struct {
	X, Y float64
}

// 方法就是一类带特殊的 接收者 参数的函数
// 接收者为结构体 Mobile
func (p Phone) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// 加法方法 接收者为结构体 Mobile
func (p Phone) Add() float64 {
	return p.X + p.Y
}

func main() {
	p := Phone{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(p.Add())
}
