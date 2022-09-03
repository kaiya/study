package main

import (
	"fmt"
	"time"
)
type myint int
type Human struct {
	name string
	age myint
	sex string
}
func (h *Human) Eat(){
	fmt.Println(h.name)
}
func say(s string){
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		println(time.RFC3339Nano)
	}
}
func main() {
	go say("world")
	say("hello")
	h := Human{"kaiya", 27, "M"}
	h.Eat()
}
