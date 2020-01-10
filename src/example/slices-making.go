package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice02("a", a)

	b := make([]int, 0, 5)
	printSlice02("b", b)

	c := b[:2]
	printSlice02("c", c)

	d := c[2:5]
	printSlice02("d", d)
}

func printSlice02(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
