package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]int)
	m["name"] = 1
	m["age"] = 24

	delete(m, "name")
	fmt.Println(m["age"])

	m["age"] = 23
	fmt.Println(m)

	v, ok := m["name"]
	fmt.Println("值为:", v, "是否在映射中:", ok)

	s := "name age"

	fields := strings.Fields(s)
	fmt.Println(fields[0])
}
