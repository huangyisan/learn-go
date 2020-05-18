package main

import "fmt"

// 如果对象实现了接口, 则他们可以被放入同一个数组中.

type Usb interface {
	// 声明两个没有实现的方式
	Start()
	Stop()
}


type Phone struct {
	Name string

}

type Camera struct {
	Name string

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


func main() {
	// 定义一个Usb类型的数组
	var a [2]Usb

	iphone := Phone{
		Name:"iphone",
	}
	camera := Camera{
		Name:"koda",
	}

	// 因为iphone和camera都实现了Usb接口,所以都可以存放到Usb的数组中
	a[0] = iphone
	a[1] = camera
	fmt.Print(a)


}
