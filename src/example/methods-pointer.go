package main

import "fmt"

type Boy struct {
	name string
	age  int
}

func (b Boy) Say() int {
	return b.age
}

func (b *Boy) Run(i int) int {
	b.age = b.age * i
	return b.age
}

func main() {
	boy := Boy{"小明", 23}
	fmt.Println(boy.age)
	boy.Run(10)
	fmt.Println(boy.Say())

}
