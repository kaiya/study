package main

import "fmt"

//返回多个值 匿名的
func foo2(a string, b int) (int, int) {
	return 1, 2
}

//返回多个值 有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	//未赋值前 r1 r2为默认值0
	r1 = 1000
	r2 = 2000
	return r1, r2
}
func main() {
	i, i2 := foo2("haha", 1)
	fmt.Println(i, i2)
	r1, r2 := foo3("a", 3)
	fmt.Println(r1, r2)

}
