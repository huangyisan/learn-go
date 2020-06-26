package main

import (
	"fmt"
)

func foo() (string) {
	var result string
	result = "ccc"
	defer func() {
		result = "Change World" // change value at the very last moment
	}()
	return result
}

func main() {
	fmt.Println("b return:", b()) // 打印结果为 b return: 2
	fmt.Println(foo())
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}