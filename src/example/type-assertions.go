package main

import "fmt"

/*
	类型断言
	判断 一个接口值是否保存了一个特定的类型
*/
func main() {
	var i interface{} = "hello.txt"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // 不会报错
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)

}
