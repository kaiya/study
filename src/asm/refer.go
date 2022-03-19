package main

import _ "unsafe"

var a = 999

func get() int

func main() {
	println(get())
	outerfunc()
}

//go:linkname gofunc
func gofunc() {
	println("func in go file")
}

func outerfunc()
