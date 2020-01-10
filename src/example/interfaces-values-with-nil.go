package main

import "fmt"

type I2 interface {
	M2()
}

type T2 struct {
	S string
}

func (t *T2) M2() {
	// 接口内的具体值为 nil 方法仍然会被 nil 接收者调用
	//为了防止触发空指针异常 在 Go 中通常会写一些方法来优雅地处理它 这里进行判断
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I2
	// t 为 nil
	var t *T2
	i = t
	describe2(i)
	i.M2()

	i = &T2{"hello.txt"}
	describe2(i)
	i.M2()
}

func describe2(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
