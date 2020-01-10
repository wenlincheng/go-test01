package main

import "fmt"

/*
	append 向切片添加元素

*/

func main() {
	var s01 []int
	printSlices01(s01)

	s01 = append(s01, 1, 2, 3)
	printSlices01(s01)

	s01 = append(s01, 2, 4, 5)
	printSlices01(s01)

}

func printSlices01(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))

}
