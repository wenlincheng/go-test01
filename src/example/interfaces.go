package main

import (
	"fmt"
	"math"
)

/*
	接口
*/

type Abser interface {
	Abs() float64
}

// 非结构体类型
type TestFloat float64

// 实现接口 Abser
func (f TestFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 结构体类型
type Computer struct {
	Weight, Price float64
}

// 实现接口 Abser
func (v *Computer) Abs() float64 {
	return math.Sqrt(v.Weight*v.Weight + v.Price*v.Price)
}

func main() {
	var a Abser
	f := TestFloat(-math.Sqrt2)
	v := Computer{3, 4}
	// 接口类型的变量可以保存任何实现了这些方法的值
	a = f // a TestFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Computer（而不是 *Computer）
	// 所以没有实现 Abser。
	// a = v 编译错误
	fmt.Println(a.Abs())
}
