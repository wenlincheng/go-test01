package main

import "fmt"

type Plain struct {
	name string
	size int
}

var maps = map[string]Plain{
	"a": {"747", 90},
	"b": {"787", 100},
}

func main() {
	fmt.Printf("%s\n", maps["a"].name)
}
