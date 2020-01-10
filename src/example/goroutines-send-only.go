package main

import "fmt"

/*
	单向信道 Send Only 、Receive Only
*/

// 唯送信道
func sendData(sendCh chan<- int) {
	sendCh <- 10
}

// 唯收信道
func receiveData(receiveCh <-chan int) {
	// 只能从信道中接收数据 会编译出错
	//receiveCh <- 10
	a := <-receiveCh
	fmt.Println(a)
}

func main() {
	// 创建一个唯送（Send Only）信道
	sendCh := make(chan<- int)
	go sendData(sendCh)
	// 只能向信道发送数据 编译出错
	// fmt.Println(<-sendCh)

	// 创建一个唯收（Receive Only）信道
	receiveCh := make(<-chan int)
	go receiveData(receiveCh)
	fmt.Println(<-receiveCh)
}
