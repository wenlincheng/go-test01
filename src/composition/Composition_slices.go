package main

import (
	"fmt"
)

// 作者
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// 博客
type post struct {
	title   string
	content string
	// 使用匿名变量嵌套结构体
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	// 嵌套一个切片 不能使用匿名切片
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Wen",
		"Lincheng",
		"Golang Coder",
	}
	post1 := post{
		"Go语言基础学习（一）——类型",
		"Go 有25个关键字或者保留字，36个预定义标识符",
		author1,
	}
	post2 := post{
		"Go语言基础学习（二）——变量",
		"Go 语言定义变量的方式很奇葩",
		author1,
	}
	w := website{
		posts: []post{post1, post2},
	}
	w.contents()
}
