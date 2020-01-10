package main

import "fmt"

// 作者
type Author struct {
	FirstName string
	LastName  string
	Bio       string
}

func (a Author) fullName() string {
	return fmt.Sprintf("%s,%s", a.FirstName, a.LastName)
}

// 博客
type Post struct {
	Title   string
	Content string
	//  匿名字段 嵌套结构体
	Author
}

func (p Post) details() {
	fmt.Println("Title: ", p.Title)
	fmt.Println("Content: ", p.Content)
	fmt.Println("Author: ", p.Author.fullName())
	fmt.Println("Bio: ", p.Author.Bio)
}

/*
	结构体嵌套 取代继承
*/
func main() {
	author1 := Author{
		"Wen",
		"Lincheng",
		"Golang Coder",
	}

	post1 := Post{
		"Go语言学习笔记",
		"Go是世界上最好的语言 Php到一边去",
		author1,
	}

	post1.details()
}
