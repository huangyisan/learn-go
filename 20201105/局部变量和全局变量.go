package main

import "fmt"

//定义在函数外部的变量为全局变量
var c int

func main() {
	//局部变量定义在{}里面的变量，只能在{}有效
	{
		i :=10
		fmt.Println(i)
	}
	// 报错
	//fmt.Println(i)

	//包括if语句中定义的变量，也是局部变量
	if a:=3; a>0 {
		fmt.Println(a)
	}
	// 报错a没有申明，a的变量只存在if中
	//a = 2

	//打印出了全局变量
	fmt.Println(c)
}
