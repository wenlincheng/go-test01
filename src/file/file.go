package main

import (
	"bufio"
	"flag"
	"fmt"
	//"github.com/gobuffalo/packr"
	"io/ioutil"
	"log"
	"os"
)

/*
	运行
	/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/bin/file -fpath=/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/src/file/hello.txt
*/
func main() {
	append()
}

/*
	绝对路径读取文件
*/
func test01() {
	data, err := ioutil.ReadFile("/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/src/file/hello.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	fmt.Println("文件内容内容：", string(data)) // 文件内容内容： Hello world!
}

/*
	使用命令行标记来传递文件路径
 	/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/bin/file -fpath=/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/src/file/hello.txt
*/
func test02() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

/*
	分块读取文件
*/
func test03() {
	//fptr := flag.String("fpath", "test.txt", "file path to read from")
	//flag.Parse()

	//f, err := os.Open(*fptr)
	f, err := os.Open("/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/src/file/hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

/*
	将文件绑定在二进制文件中
*/
func test04() {
	//box := packr.NewBox("../filehandling")
	//data := box.String("test.txt")
	//fmt.Println("Contents of file:", data)
}

// 将字符串写入文件
func createString() {
	f, err := os.Create("/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/src/file/string.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 将字节写入文件
func writeByte() {
	f, err := os.Create("/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/src/file/bytes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 将字符串一行行写入文件
func writeLine() {
	f, err := os.Create("/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/src/file/lines.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.",
		"It is easy to learn Go."}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

// 追加到文件
func append() {
	f, err := os.OpenFile("/Users/wenlincheng/go/src/git.daddylab.com/DaddyLab/go-test01/src/file/lines.txt",
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
