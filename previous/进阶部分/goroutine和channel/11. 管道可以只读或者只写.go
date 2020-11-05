package main

import "fmt"

// 默认为可读可写。

func main() {
	// 声明为只写 <-
	var chan2 chan <- int
	chan2 = make(chan int, 10)
	chan2 <- 1
	// 如果读取则报错
	// num := <- chan2


	// 声明为只读
	var chan3 <-chan int
	num2 := <- chan3
	// 如果写入就报错
	//chan3 <- 1
	fmt.Println(num2)
}
