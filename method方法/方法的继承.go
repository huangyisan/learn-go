package main


//method是可以继承的，如果匿名字段实现了一个method，那么包含该匿名字段的struct也能调用该method

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  //匿名字段
	school string
}
type Employee struct {
	Human   //匿名字段
	company string
}


// 匿名字段实现了SayHi的函数，所以所有存在Human的struct都有SayHi的方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	//mark是基于Student进行实例化，Student存在Human这个匿名字段，所以可以使用SayHi这个方法。 sam类似
	mark.SayHi()
	sam.SayHi()
}
