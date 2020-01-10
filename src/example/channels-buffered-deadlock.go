package main

import (
	"fmt"
)

/*
	缓冲信道死锁  deadlock
*/
func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	// 由于创建的信道容量为2 已经向信道中发送了两个数据 再添加会超出信道的容量 因此会出现阻塞
	// 程序中没有其他协程来读取信道中的数据 因此会出现死锁 触发 panic
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
