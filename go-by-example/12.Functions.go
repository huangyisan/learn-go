package main

import "fmt"

// 将函数作为参数传入，先定义type
type addfunc func(int,int) int

func add(a,b int) int {
	return a + b
}


func myAdd(f addfunc) int {
	c := f(1,2)
	return c
}



func main() {
	a := add
	res := myAdd(a)
	fmt.Println(res)
}