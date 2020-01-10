package main

import "fmt"

func main() {
	primes := [8]int{1, 2, 2, 4, 5, 6, 7, 8}
	s := primes[2:4]
	fmt.Println(s)
}
