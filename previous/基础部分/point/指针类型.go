package main

import "fmt"
func main() {
	var i int = 10
//	 i的地址为 &i
fmt.Println("i的地址为=",&i)

	// 指针类型, 指针变量存储的是一个地址, 这个地址指向的内存空间才是值
	// *变量 指针对应的是变量内存的地址&xxx, *和&对应
	var ptr *int = &i
	fmt.Println(*ptr)



}