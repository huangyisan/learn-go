package main

import "fmt"

// 定义了一个speak接口，有sayHi()的方法
type speak interface{
	sayHi()
}


type student struct {

}

type teacher struct {

}

// student实现了sayHi()方法
func (s student) sayHi() {
	fmt.Println("我是学生")
}

// teacher实现了sayHi()方法
func (s teacher) sayHi() {
	fmt.Println("我是老师")
}

// 函数的参数为接口类型，此时传入不同的参数，得到不同的sayHi结果
func whoSay(i speak) {
	i.sayHi()
}


func main() {

	var s student
	var t teacher

	// 函数的参数为接口类型，此时传入不同的参数，得到不同的sayHi结果
	whoSay(s)
	whoSay(t)

	// 接口放入切片中，进行遍历
	l := []speak{s,t}

	for _,i := range l{
		i.sayHi()
	}

}