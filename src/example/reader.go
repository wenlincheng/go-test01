package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 读取为流
	r := strings.NewReader("Hello, Reader!")
	// 每次读取8个字节
	b := make([]byte, 8)
	for {
		// 返回填充的字节数和错误值
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// 流的结尾判断
		if err == io.EOF {
			break
		}
	}
}
