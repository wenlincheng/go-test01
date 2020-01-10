package main

import "fmt"

type Teacher struct {
	Name string
	age  int
}

/*
	make 与 new的区别
*/
func main() {
	// 返回类型本身 分配内存空间 并进行初始化
	i := make([]int, 4)
	fmt.Println(i) // [0 0 0 0]
	m := make(map[string]string, 9)
	fmt.Println(m) // map[]

	// 返回类型的指针 分配内存空间不进行初始化
	j := new([]int)
	fmt.Println(j) // &[]
	t := new(Teacher)
	fmt.Println(t) // &{ 0}
}
