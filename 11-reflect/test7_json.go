package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"xiju", 1999, 98, []string{"xingye", "zhangbozhi"}}

	// 编码的过程 把结构体转换成JSON
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)
	//解码的过程,把JSON字符串解码为结构体
	// jsonStr = {"title":"xiju","year":1999,"rmb":98,"actors":["xingye","zhangbozhi"]}

	myMovie := Movie{}

	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", myMovie)

}
