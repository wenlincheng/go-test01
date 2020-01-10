package main

import "fmt"

type Animal struct {
	height int
	weight int
}

func main() {
	a := Animal{1, 1000}
	fmt.Println(a.height)
	a.height = 2
	fmt.Println(a.height)
}
