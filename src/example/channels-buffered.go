package main

import "fmt"

/*
	带缓冲的信道
*/
func main() {
	// 修改为1 会报错 fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
