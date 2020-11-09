package main

import "fmt"

// 只要两个接受者不一样，方法就不会被重载，也就是即便是方法名一样，不同的接受者会有不同的行为。



type myInt int

type str string

func (num myInt) myPrint() {
	fmt.Println(num)
}

func (s str) myPrint() {
	fmt.Println(s)
}

func main() {
	var s str
	var age myInt
	s = "yisan"
	age = 29

	// 虽然myPrint存在两个同名方法，但作用的对象不同。不会出现方法重复定义的错误。
	s.myPrint()
	age.myPrint()
}