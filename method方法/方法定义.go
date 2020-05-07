package main

import (
	"fmt"
	"math"
)

// go中没有类的概念
// 在方法名称之前给定一个receiver,比如下面， t是receiver， methodName为方法
// receiver不能是一个接口类型，不能是一个指针类型，但是它可以是任何其他允许类型的指针。
// func (t Type) methodName(parameter list)(return list) {}
// https://zhuanlan.zhihu.com/p/78898099

type Circle struct {
	r float64
	pi float64
}

// 非结构体，则需要type关键字进行 类型等价定义。否则不能使用方法。func (s int) area() {} 这种写法是错误的，因为int没有进行类型等价定义。
type Myslice []int

// receiver是Circle, 是结构体
func (s Circle) area() float64 {
	squire := math.Pow(s.r,2) * s.pi
	return squire
}

func (s Myslice) printRange() {
	for _,v :=range s{
		fmt.Printf("%v\n",v)
	}

}

func main() {
	// 类似面向对象的实例化对象
	myCircle := Circle{
		1.5,
		3.14,
	}
	// 类似调用实例化对象的函数，就成了方法
	fmt.Printf("圆面积为%f\n", myCircle.area())

	// 类似实例化myslice
	myslice := Myslice{1,2,3,4,5}

	myslice.printRange()
}

