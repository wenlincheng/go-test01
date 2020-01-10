package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(a)

	// 截取切片长度为0
	a = a[:0]
	printSlice(a)

	// 拓展切片长度
	a = a[:4]
	printSlice(a)

	// 舍去前两个值
	a = a[2:]
	printSlice(a)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
