package main

import "fmt"

//声明全局变量
var gA int

func main() {
	var a int // default 0

	var b int = 100
	//初始化时省去数据类型 通过值自动匹配当前变量的数据类型
	var c = 100

	fmt.Printf("type of c = %T\n", c) //d s v T
	//省去var关键字 直接自动匹配(只能用在函数体内 不能用于全局变量)
	e := 100
	var xx, yy int = 100, 200
	var kk, ll = 100, "kaiya"
	var (
		vv int  = 100
		jj bool = true
	)
	_, _, _, _, _, _, _, _, _ = a, b, e, xx, yy, kk, ll, vv, jj
}
