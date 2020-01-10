package main

import "fmt"

/*
	延迟执行应用

*/

func main() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	// 字符串反转
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
	defer fmt.Printf("\nReversed String: ")

	// 数组数据反转
	nums := []int64{1, 2, 3, 4}
	for _, num := range nums {
		defer fmt.Printf("%d", num)
	}
	defer fmt.Printf("Reversed num:")

}
