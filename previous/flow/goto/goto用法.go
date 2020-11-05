package main

import "fmt"

// 直接跳到label对应的语句进行执行
// 建议少用


func main() {
	fmt.Println("1")
	// 直接跳到abc
	goto abc
	fmt.Println("2")
	fmt.Println("3")
	abc:
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
}