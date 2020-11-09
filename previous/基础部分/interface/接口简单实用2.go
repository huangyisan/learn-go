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


func main() {
	// 申明speaker
	var speaker speak

	var s student
	var t teacher

	// 将s赋值给speaker
	speaker = s
	// 此时的speaker便是student
	speaker.sayHi()

	// 将t赋值给speaker
	speaker = t
	// 此时的speaker便是teacher
	speaker.sayHi()
}