package main

import "fmt"

type Bed struct {
	X, Y float64
}

// 方法
func (b *Bed) Scale(f float64) {
	b.X = b.X * f
	b.Y = b.Y * f
}

// 函数
func ScaleFunc(b *Bed, f float64) {
	b.X = b.X * f
	b.Y = b.Y * f
}

func main() {

	b := &Bed{1.9, 1.3}
	// bed 为值不是指针，带指针接收者的方法也能直接被调用
	// go 会将语句 bed.Scale(2.0) 解释为(&bed).Scale(2.0)
	b.Scale(2.0)
	fmt.Println(b)
	ScaleFunc(b, 2.0)
	fmt.Println(b)

	bed := Bed{1.9, 1.3}
	// bed 为值不是指针，带指针接收者的方法也能直接被调用
	// go 会将语句 bed.Scale(2.0) 解释为(&bed).Scale(2.0)
	bed.Scale(2.0)
	fmt.Println(bed)
	//ScaleFunc(bed, 2.0) // 编译错误
	fmt.Println(bed)
}
