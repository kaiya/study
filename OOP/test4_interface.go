package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleeping")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
}
func main() {
	var animal AnimalIF //接口的数据类型 父类指针

	animal = &Cat{"Green"}
	animal.Sleep() //调用的就是Cat的Sleep方法

}
