package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func computeInt(fn func(int, int) int, x, y int) int {
	return fn(x, y)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	multiply := func(x, y int) int {
		return x * y
	}
	// 函数值作为参数传递
	fmt.Println(computeInt(multiply, 2, 2))
	// 函数值
	fmt.Println(hypot(5, 12))
	// 函数值作为参数传递
	fmt.Println(compute(hypot))
	// 求以 x 为底的 y 次方
	fmt.Println(compute(math.Pow))

}
