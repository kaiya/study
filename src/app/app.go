package main

import (
	"awesomeProject/src/greet"
	"fmt"
	"runtime"
)

func main() {
	println(runtime.Version())

	fmt.Println(greet.Version)
}
