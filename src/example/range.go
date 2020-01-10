package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	for i, e := range a {
		fmt.Printf("%d=%d\n", i, e)
	}

	for i := range a {
		fmt.Printf("%d\n", i)
	}

	for _, e := range a {
		fmt.Printf("%d\n", e)
	}

}
