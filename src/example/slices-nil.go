package main

import "fmt"

func main() {
	// nil 切片的长度和容量为 0 且没有底层数组
	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("a is nil")
	}
}
