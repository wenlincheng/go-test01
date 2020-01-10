package main

import "fmt"

type I4 interface {
}

func main() {
	var i I4
	// 或者直接使用空接口 var i interface{}
	describe4(i)

	i = 42
	describe4(i)

	i = "hello.txt"
	describe4(i)
}

func describe4(i I4) {
	fmt.Printf("(%v, %T)\n", i, i)
}
