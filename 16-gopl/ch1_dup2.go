package main

import (
	"bufio"
	"fmt"
	"os"
)

type Human struct {
	name string
	age  int
}

func main() {
	p := new(int) //new(T)创建一个未命名的T类型变量初始化为T类型的零值,并返回其地址
	fmt.Println(*p)
	//gopl.io/ch1/dup2
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
