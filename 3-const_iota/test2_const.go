package main

//const来定义枚举类型
const (
	//可以在const()添加一个关键字 iota, 每行的iota都会累加1 第一行的iota的默认值是0
	BEIJING   = iota //iota = 0
	SHANGHAI         //iota = 1
	GUANGZHOU        //iota = 2
)

//iota只能配合const()一起使用,iota只有在const进行累加效果
const (
	//类似于为之后的定义好一个公式
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	//常量 只读
	const length int = 10

}
