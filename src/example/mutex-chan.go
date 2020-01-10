package main

import (
	"fmt"
	"sync"
)

var z = 0

func increment3(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<-ch
	wg.Done()
}

/*
	信道解决并发操作共享资源的问题
*/
func main() {
	var w sync.WaitGroup
	// 由于缓冲信道的容量为 1，所以任何其他协程试图写入该信道时，都会发生阻塞，直到 x 增加后，信道的值才会被读取
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment3(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of z", z)
}
