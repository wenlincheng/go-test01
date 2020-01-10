package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

/*
	信道
*/
func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// 计算前三个的和
	go sum(s[:len(s)/2], c)
	// 计算后三个的和
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
