package main

import "fmt"

/*
	信道转换
*/

func sendData2(sendCh chan<- int) {
	sendCh <- 10
	// 唯送信道 无法从信道中读取数据 编译出错
	// fmt.Println(<-sendCh)
}

func main() {
	cha1 := make(chan int)
	// 转换成唯送信道
	go sendData2(cha1)
	// 双向信道 编译没问题
	fmt.Println(<-cha1)
}
