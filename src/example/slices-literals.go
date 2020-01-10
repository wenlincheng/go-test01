package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)

	b := []bool{true, false, true, false}
	fmt.Println(b)

	c := []struct {
		x int
		y string
	}{
		{20, "小明"},
		{23, "小红"},
	}
	fmt.Println(c)
	fmt.Println(c[0].x)
}
