package main

import "fmt"

// 非匿名函数用完后，里面的变量就会释放
func test01() int {
	var x int
	x+=1
	return x
}

// 定义一个函数，返回一个匿名函数，这里要看清后面，func() int是一个整体，表示一个返回类型为int的匿名函数
func test02() func() int {
	var x int

	return func() int {
		//在匿名函数中对x进行了修改
		//只要该匿名函数一直被使用，则内部的变量会被一直修改。
		x+=1
		return x
	}
}

func main() {
	a := test01
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())


	b := test02()
	//此时对b调用，则为调用了匿名函数
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}