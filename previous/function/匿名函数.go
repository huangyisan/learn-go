package main

import "fmt"

// 全局变量
var (global = func(n1 int,n2 int) int {
	return n1 + n2}
)

func main() {
	// 在定义匿名函数时候直接调用，这种方式匿名函数只能被调用一次。
	// 匿名函数不需要写函数名
	res1 := func(n1 int,n2 int) int {
		return n1 + n2
	//	直接跟随()进行调用
	}(10, 20)
	fmt.Println(res1)

//	 匿名函数赋值给变量，则该变量为函数类型

	res := func(n1 int,n2 int) int {
		return n1 + n2
		//	直接跟随()进行调用
	}
	fmt.Printf("变量res类型为%T\n", res)
	fmt.Printf("传入参数后执行%d", res(19,10))

//	 全局匿名函数
	res2 :=global(9,10)
	fmt.Println(res2)
}