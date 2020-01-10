package main

import "fmt"

type Car struct {
	price float32
	brand string
}

func main() {
	c := Car{20000.0, "宝马"}
	p := &c
	fmt.Println(p.price)
	fmt.Println(c.brand)
}
