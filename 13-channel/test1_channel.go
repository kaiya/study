package main

import "fmt"

func main() {
	//定义一个channel

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 运行结束")
		fmt.Println("goroutine 运行中...")

		c <- 666 //将666发送给c
	}()
	num := <-c //在main goroutine中 从c中接收数据,并赋值给num
	//channel有同步main和goroutine的能力 所以每次都是go routine先结束 即使子协程还没往channel写数据, main协程会阻塞等待
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束")
}
