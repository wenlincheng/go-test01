package main

import "fmt"

/*
	nil 接口值
*/
type I3 interface {
	M3()
}

func main() {
	var i I
	describe3(i)
	// 由于接口值为nil 编译错误
	// i.M3()
}

func describe3(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
