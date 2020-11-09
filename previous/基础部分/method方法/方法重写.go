package main

import "fmt"

//方法是可以继承和重写的
//存在继承关系时，按照就近原则，进行调用


type children struct {
	name string
	age int
}

type boy struct {
	children
	class string
}

// 定义Human的方法
func (s children) printName() {
	fmt.Println(s.name)
}

// 在Student中对printName方法重写
func (s boy) printName() {
	fmt.Println(s.name + "X")
}

func main() {
	Lilei := boy{children{name:"lilei", age:19}, "one"}

	//Lilei从boy的struct实例化而来，根据就近原则，使用boy struct的printName方法。
	Lilei.printName()  // lileiX

}