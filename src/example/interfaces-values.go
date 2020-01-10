package main

import (
	"fmt"
	"math"
)

// 接口值

// 定义接口
type I interface {
	M()
}

// 定义结构体
type T struct {
	S string
}

// 实现接口
func (t *T) M() {
	fmt.Println(t.S)
}

// 定义非结构体类型
type F float64

// 实现方法
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	// 接口值 保存了包含的值和具体类型
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
