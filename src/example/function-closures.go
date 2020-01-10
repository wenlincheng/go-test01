package main

import "fmt"

/*
	函数闭包
*/
func adder() func(int) int {
	sum := 0
	// 函数值
	fun02 := func(x int) int {
		sum += x
		return sum
	}
	return fun02
}

// 返回一个斐波那契数列
func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		z := x
		x = y
		y = x + z
		return z
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 1; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	// 斐波那契数列
	// F(1)=1，F(2)=1, F(n)=F(n-1)+F(n-2)（n>=3，n∈N*）
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
