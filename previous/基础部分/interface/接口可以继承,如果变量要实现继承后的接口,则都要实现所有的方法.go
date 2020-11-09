package main

import "fmt"

type interface1 interface {
	t1()
}


type interface2 interface {
	t2()
}

type All interface {
	// all继承了interface1和interface2的接口
	interface1
	interface2
	t3()
}

type Human1 struct {

}

func (s Human1) t1()  {
	fmt.Println("t1")
}

func (s Human1) t2()  {
	fmt.Println("t2")
}

func (s Human1) t3()  {
	fmt.Println("t3")
}

type test struct {

}

func (s test) c(all All) {
	all.t1()
	all.t2()
	all.t3()
}
func main() {
	p1:=Human1{}
	h:=test{}

	// 如果少实现一个方法,则无法执行h.c(p1)
	h.c(p1)

}