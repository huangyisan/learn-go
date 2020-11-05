package main

import "fmt"

//1. new: 用来分配内存，主要用来分配值类型，比如int,float32,struct 返回的是指针
//2. make: 用来分配内存， 主要用来分配引用类型， 比如channel、map、slice。


func main() {
	num1 := 100
	fmt.Printf("num1类型%T, num1的值=%v, num1的地址%v\n", num1, num1, &num1)
	//num1 的类型为int， 值为空间存放的100， 内存地址表示该空间的内存地址

	num2 := new(int)
	fmt.Printf("num2类型%T, num2的值=%v, num2的地址%v，num2的值指向的值为%v\n", num2, num2, &num2, *num2)
	//	num2的类型为指针，值为空间存放的内存地址， 指针地址为该存放了内存地址的空间内存地址， 空间存放的内存地址指向的值为零值，因为这里是int,所以该值为0

	//此时num2为指针类型。 给他赋值
    *num2 = 100
    fmt.Printf("重新赋值后的值为%v", *num2)

}