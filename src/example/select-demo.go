package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1000 * time.Millisecond)
	// 会阻塞 直到某个分支可以继续执行 由于server2休眠3秒 而server1休眠时间为
	// 为6秒 因此，主协程阻塞3秒 待server2向信道中发送数据后 主协程继续执行
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

}
