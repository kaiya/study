package main

import (
	"fmt"
	"runtime"

	"github.com/Kaiya/golangStudy/src/greet"
)

func main() {
	println(runtime.Version())

	fmt.Println(greet.Version)
}
