package main

import (
	"fmt"
	"time"
)

//子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

//main goroutine
//子goroutine依赖于main goroutine
func main() {
	// 创建一个go程去执行newTask
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine : i  = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
