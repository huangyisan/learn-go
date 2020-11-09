package main

import "fmt"

//回调函数，函数的参数是函数类型，这个函数就是回调函数。
//go中多态的实现

// 假定一个sum函数
func Sum(a,b int) int {
	return a+b
}

// 定义一个MyFunc函数，为func(int,int) int函数格式
// func(int,int) int格式和Sum函数一致。实现了多态
type MyFunc func(int,int) int

// 定义一个Count函数，参数第一个对象为函数，后面跟随两个int
func Count(myfunc MyFunc, a,b int) {
	var sum int
	sum = myfunc(a,b)
	fmt.Println(sum)
}

func main() {
	//Count中将Sum函数作为参数传入，因为Sum函数也是func(int,int) int函数格式
	Count(Sum,1,2)
}