package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) // 注释掉休眠时间 主协程退出 协程运行也会停止
		fmt.Println(s)
	}
}

/*
	Go 协程（goroutine）是由 Go 运行时管理的轻量级线程
*/
func main() {
	go say("world")
	say("hello.txt")
}
