package main

import "fmt"

// AddUpper()是一个函数，返回类型为 func (int) int
func AddUpper() func(int) int {
	// 闭包内容 ，返回的是一个匿名函数， 但是这个匿名函数引用到了函数外的n，因此这个匿名函数就和n形成一个引用的整体关系，构成闭包

	var n int = 10
	return func(x int) int {
		n = n+x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

// 可以理解为闭包是类, 函数是操作,n是字段，函数和他使用到的n构成闭包
// 当我们反复调用f函数的时候，因为n是初始化一次，因此每调用一次就进行累计
// 要搞清楚闭包的关键，就是要分析出返回的函数他使用(引用)了哪些变量，因为函数和它引用到的变量共同构成闭包。
