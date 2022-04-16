package lib1

import "fmt"

func init() {
	fmt.Println("lib1.init()....")
}

//当前lib1包提供的接口
func Lib1Test() {
	fmt.Println("lib1Test()...")
}
