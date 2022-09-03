package main

import (
	_ "awesomeProject/src/5-init/lib1"
	//mylib2 "awesomeProject/src/5-init/lib2"
  . "awesomeProject/src/5-init/lib2"
)

func main() {
	//lib1.Lib1Test()
	//mylib2.Lib2Test()
  Lib2Test()
}
