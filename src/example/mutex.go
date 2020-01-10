package main

import (
	"fmt"
	"sync"
)

var y = 0

func increment2(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

/*
	mutex解决并发操作共享资源的问题
*/
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of y", y)
}
