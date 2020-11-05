package main

import "fmt"

// 数组类型并不是引用类型，每次调用都是会重新开辟内存区。

// 可以使用指针的方式修改调用的对象

func zero(ptr *[2]int){

	*ptr = [2]int{3,4}
}

func main() {
	c := [2]int{1,2}
	// 传入内存地址。内存地址的*就是原数据
	zero(&c)
	fmt.Println(c)
}