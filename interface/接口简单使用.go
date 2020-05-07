package main

import (
	"fmt"
)

//go的多态通过接口实现


//1. 接口类型只能定义一组方法,且方法不需要在接口定义中实现.
//2. 接口不能包含任何变量.
//3. 不需要显示的声明某个结构体或者变量实现了接口,只要该结构体或者变量实现了接口定义的全部方法,则他们就实现了这个接口.
//4. interface是引用类型,零值为nil

//5. 空接口没有任何方法,所以所有类型都实现了空接口.

// 基本语法
//type 接口名 interface {
//	method1()
//	method2()
//	method3()
//	method4()
//}

//声明一个接口

type Usb interface {
	// 声明两个没有实现的方式
	Start()
	Stop()
}


type Phone struct {

}

type Camera struct {

}

// 让Phone实现接口的两个方法

func (s Phone) Start() {
	fmt.Println("手机开始工作")
}


func (s Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让camera实现接口的两个方法
func (s Camera) Start() {
	fmt.Println("手机开始工作")
}


func (s Camera) Stop() {
	fmt.Println("手机停止工作")
}


// 计算机
type PC struct {

}

//	接受usb接口类型,只要实现了usb接口的方法,就是指实现了usb接口声明的所有方法
func (s PC) working(usb Usb) {
	// 通过usb接口变量来调用start和stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	// 创建结构体
	c:=PC{}
	iphone := Phone{}
	camera := Camera{}

	//将对象传入,展现了多态的情形
	c.working(iphone)
	c.working(camera)
}