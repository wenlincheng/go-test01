package main

import "fmt"

// 结构体
type Student struct {
	name string
	age  int
}

func main() {
	s := Student{"小红", 23}
	fmt.Println(s.name)
	fmt.Println(s.age)
}
