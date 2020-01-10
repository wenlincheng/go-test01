package main

import "fmt"

func main() {
	names := [6]string{
		"Tom",
		"Jevin",
		"Paul",
		"Ringo",
	}

	fmt.Println(names)

	a := names[1:3]
	b := names[1:2]
	fmt.Println(a, b)
	b[0] = "ZYX"
	fmt.Println(a, b)
	fmt.Println(names)

}
