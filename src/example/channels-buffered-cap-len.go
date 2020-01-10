package main

import (
	"fmt"
)

/*
	缓冲信道的容量、长度
	容量为创建信道的时候的容量
	长度是信道中当前排队元素的个数
*/
func main() {
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))   // 3
	fmt.Println("length is", len(ch))     // 2
	fmt.Println("read value", <-ch)       // naveen
	fmt.Println("new length is", len(ch)) // 1
}
