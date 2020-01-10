package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 6, 8, 9, 41}

	s = s[1:5]
	fmt.Println(s)

	s = s[:5]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
