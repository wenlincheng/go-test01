package main

import (
	"fmt"
	"math"
)

// 非结构体类型
type MyFloat float64

type MyInt int

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (i MyInt) Add() int {
	return int(i) + int(i)
}

func main() {
	// 2 的平方根
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// 1
	i := MyInt(1)
	fmt.Println(i.Add())
}
