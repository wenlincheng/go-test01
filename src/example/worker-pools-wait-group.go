package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

/*
	WaitGroup 使用计数器来工作。当我们调用 WaitGroup 的 Add 并传递一个 int 时，
	WaitGroup 的计数器会加上 Add 的传参。
	要减少计数器，可以调用 WaitGroup 的 Done() 方法。
	Wait() 方法会阻塞调用它的 Go 协程（主协程），直到计数器变为 0 后才会停止阻塞。
*/
func main() {
	// 3个go协程
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	// 主协程等待计数器变为0
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
