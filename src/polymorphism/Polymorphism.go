package main

import (
	"fmt"
)

/*

	多态

*/

// 接口
type Animal interface {
	weight() int64
	run() int64
}

// 袋鼠
type Kangaroo struct {
	Name    string
	FeetNum int64
	Weight  int64
}

// 狗
type Dog struct {
	Name    string
	FeetNum int64
	Weight  int64
}

func (kangaroo Kangaroo) weight() int64 {
	return kangaroo.Weight
}

func (kangaroo Kangaroo) run() int64 {
	return kangaroo.Weight
}

func (dog Dog) weight() int64 {
	return dog.Weight
}

func (dog Dog) run() int64 {
	return dog.Weight
}

func calculateNetIncome(animals []Animal) {
	var totalWeight int64
	for _, animal := range animals {
		fmt.Printf("重量为%dkg，%d条腿\n", animal.weight(), animal.run())
		totalWeight += animal.weight()
	}
	fmt.Printf("总重量为：%d", totalWeight)
}

func main() {
	animal1 := Kangaroo{Name: "Kangaroo", Weight: 20, FeetNum: 2}
	animal2 := Dog{Name: "Dog", Weight: 20, FeetNum: 4}
	incomeStreams := []Animal{animal1, animal2}
	calculateNetIncome(incomeStreams)
}
