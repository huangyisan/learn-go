package main

import (
	"fmt"
	//不同级目录，需要引入包名
	"diff"
)

func main()  {
	//1. 同一个目录，调用别的文件的函数，不需要引入包名，可以直接调用。
	a := hello()
	fmt.Println("main函数")
	fmt.Println(a)
	diff.Diff()
}
