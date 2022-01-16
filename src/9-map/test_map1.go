package main

import "fmt"

func main() {
	//声明myMap1是map类型，key是string，Value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("nil map")
	}
	//在使用map前用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "C++"
	myMap1["three"] = "Golang"

	fmt.Println(myMap1)
	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "Golang"
	fmt.Println(myMap2)
	//第三种
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "C++",
		"three": "Golang",
	}
	fmt.Println(myMap3)

	cityMap := make(map[string]string)
	//添加
	cityMap["China"] = "Beijing"
	//遍历
	printMap(cityMap)
	//删除
	delete(cityMap, "China")
	//修改
	cityMap["China"] = "Guangzhou"

}

func printMap(cityMap map[string]string) {
	//citymap是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}
