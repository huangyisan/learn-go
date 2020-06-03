package main

import "fmt"

// 常量const定义,且必须定义的时候初始化
// 常量不能修改
// 常量只能bool, 数值类型(int float系列), string类型.


func main() {
	// 常量用const定义,且必须初始化
	const a = 1
	fmt.Println(a)

	// iota自增
	const (
		b = iota + 1 // iota表示0,仅在第一次const定义的()范围内
		c
		d
	)
	fmt.Println(b,c,d)

	const (
		// 多次定义, 则第一个位置的iota为0, 第二个位置的iota为1, 第三个位置的iota为2. 改变任何iota的初始位置不会影响到其他的被定义了的iota的常量的值.
		e = iota + 1
		f = iota
		g = iota
	)
	fmt.Println(e,f,g)
}
