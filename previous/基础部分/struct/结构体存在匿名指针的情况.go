package main

import "fmt"

type persion struct {
	name string
}

type student struct {
	int
	*persion
}

func main() {
	//第一种方法，直接赋值，里面使用&的方式进行赋值处理。
	s := student{
		29,
		&persion{"yisan"},
	}

	// 第二种方法，进行new分配空间(用new分配空间实质是开辟内存，并返回指针)
	var s2 student
	// 使用new开辟内存，并返回指针
	s2.persion = new(persion)
	s2.persion.name = "ys"
	s2.int = 28

	fmt.Println(s.name, s.int)
	fmt.Println(s2.name, s2.int)
}