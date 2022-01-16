package lib2

import "fmt"

func init() {
	fmt.Println("lib2.init()....")
}

//当前lib2包提供的接口
func Lib2Test() {
	fmt.Println("lib2Test()...")
}
