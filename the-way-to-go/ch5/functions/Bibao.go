package functions

import "fmt"

// squares的例子证明，函数值不仅仅是一串代码，还记录了状态。在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。这就是函数值属于引用类型和函数值不可比较的原因。Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。

// 返回一个匿名函数
func square() func() int {
	var x int
	return func() int {
		// x突破了变量域, 可以访问到上层函数的x
		x++
		return x * x
	}
}

func CallSquare() {
	fmt.Println(square()())
	fmt.Println(square()())
	fmt.Println(square()())
}
