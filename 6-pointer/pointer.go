package main

import "fmt"

func swap(pa *int, pb *int) {
	var tmp int
	tmp = *pa //tmp = main::a
	*pa = *pb //main::a = main::b
	*pb = tmp //main::b = temp
}
func main() {
	var a int = 10
	var b int = 20
	//swap
	swap(&a, &b)
	fmt.Println("a= ", a, " b= ", b)

	var p *int = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int = &p //二级指针
	fmt.Println(&p)
	fmt.Println(pp)
}
