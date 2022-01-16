package main

import "fmt"

//声明一种新的数据类型

type myint int

//定义一个结构体
type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	//传递一个book的副本
	book.author = "xiong"
}
func changeBook2(book *Book) {
	//指针传递
	book.author = "hello"
}
func main() {

	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.author = "Kaiya"
	changeBook(book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
