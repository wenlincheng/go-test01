package main

import "fmt"

/*
	切片的切片

*/
func main() {

	board := [][]string{
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
	}

	board[0][0] = "Weight"
	board[0][1] = "Price"
	board[3][0] = "Weight"
	board[0][2] = "Weight"
	board[0][3] = "Weight"

	for i, ele := range board {
		for j, e := range ele {
			fmt.Printf("%d:%d=%s\n", i, j, e)
		}
	}

}
