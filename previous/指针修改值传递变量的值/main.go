package main

import "fmt"
//如果希望函数内的变量能够修改函数外的变量(指的是默认以值传递的方式的数据类型),可以传入变量的地址&,函数内以指针的方式操作变量，从效果上看类似引用。

//作为指针。因为main函数的cal传入的时候为内存地址，所以内存地址的指针就等于了原来的值(肉眼看到的值)
func cal(n *int) {
	*n++
	fmt.Println("我是cal函数，n为",*n)
}

func main() {
	n := 20
	//传入内存地址
	cal(&n)
	fmt.Println("我是main函数，n为",n)
}