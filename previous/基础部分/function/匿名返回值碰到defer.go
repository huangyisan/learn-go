package main

import (
	"fmt"
)

func main() {
	fmt.Println("a return:", a()) // 打印结果为 a return: 0
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}