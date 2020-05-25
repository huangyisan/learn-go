package main

import "fmt"

// 数组定义
// var 数组名 [长度]数组类型

func main()  {
	var intArry [3]int
	// 数组名称的内存地址指向该数组收个元素的内存地址
	fmt.Printf("intArry地址为%p\n", &intArry)
	fmt.Printf("数组首地址为%p", &intArry[0])

}
