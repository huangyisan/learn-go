package main

import "fmt"

// 函数也是一种数据类型。可以赋值给一个变量，该变量则为函数类型的变量，通过该变量可以对函数调用。

func getsum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型，可以作为形参直接传入
func myFunc(funcvar func(int, int) int, n1 int, n2 int) int {
	return funcvar(n1,n2)
}

func main() {
	a := getsum
	fmt.Printf("a的数据类型为%T, getsum数据类型为%T", a, getsum)

	b1 := 1
	b2 := 2

	b := myFunc(getsum,b1,b2)

	fmt.Println(b)
}
