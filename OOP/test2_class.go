package main

import "fmt"

//类名首字母大写 表示其他包也可以访问
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}
func (this Hero) GetName() string {
	return this.Name
}
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

// func (this Hero) SetName(newName string) {
// 	//this是调用该方法的对象的一个副本（拷贝）
// 	this.Name = newName
// }
func main() {
	hero := Hero{Name: "zhangsan", Ad: 100, Level: 1}
	hero.show()
	hero.SetName("li4")
	hero.show()
}
