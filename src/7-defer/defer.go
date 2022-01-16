package main

import "fmt"

func main() {
	defer fmt.Println("main end1")
	defer fmt.Println("main end2") //因为是压的方法栈 所以end2 先执行
	fmt.Println("hello go 1")
	fmt.Println("hello go 2")
	//!!defer在return执行后执行
}
