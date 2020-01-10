package main

import (
	"fmt"
	"time"
)

func process2(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

/*
	在没有 case 准备就绪时，可以执行 select 语句中的默认情况（Default Case）
	这通常用于防止 select 语句一直阻塞
*/
func main() {
	ch := make(chan string)
	go process2(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}
