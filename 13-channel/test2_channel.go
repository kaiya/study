package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带有缓冲的channel
	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	go func() {
		defer fmt.Println("sub goroutine end")
		//前3个不会阻塞 第4个阻塞等待 被取走
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("sub goroutine running, sending elem:", i, " len(c)=", len(c), " cap(c)=", cap(c))

		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c //从c中接收 并赋值给num, c和-之间没有空格
		fmt.Println("num = ", num)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("main end")
}
