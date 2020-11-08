package main

import "fmt"

type myString string

type student struct {
	name string
	int
	myString
}

func main() {
	var s student
	s.name = "yisan"
	//直接使用变量.类型的方式进行赋值。
	s.int = 29
	s.myString = "ok"

	fmt.Println(s.name, s.int, s.myString)
}