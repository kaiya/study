package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组，切片slice

	fmt.Printf("myArray type is %T\n", myArray)
	printArray(myArray)
	fmt.Println("=======")
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}

//slice传递的是指针
func printArray(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}
