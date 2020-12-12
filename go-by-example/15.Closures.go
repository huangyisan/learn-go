package main

import "fmt"

// 当函数没结束的时候，匿名函数的空间一直会保留，所以多次调用闭包，a的值会累加
func bibao() func() int {
	a := 1
	return func() int {
		a += 1
		return a
	}
}

func main() {
	b := bibao()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
