package main

import "fmt"



func main() {
	var a = 3
	var i int

	// 单个条件循环
	for i < a {
		fmt.Println(i)
		i ++
	}

	// 经典循环
	for i = 1; i< a; i++ {
		fmt.Println(i)
	}

	// 等同python while
	for {
		fmt.Println("无限循环")
		break
	}
}