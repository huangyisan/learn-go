package main

import "fmt"

// 同级目录被调用，函数首字母无需大写。
func hello() int {
	fmt.Println("hello")
	return 1
}
