package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
	//interface{} 该如何区分 此时引用的底层数据类型到底是什么呢
	//给interface{}提供了一种类型断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type,value = ", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"golang"}
	myFunc(book)
	myFunc("hello")
}
