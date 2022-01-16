package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//用go创建一个形参为空,返回值为空的函数
	go func() {

		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	//如何在main goroutine中获取子goroutine的结果bool呢? 用到了channel
	go func(a int, b int) bool {

		fmt.Println("a = ", a, " b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
