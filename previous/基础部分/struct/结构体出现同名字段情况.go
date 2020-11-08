package main

import "fmt"

// 当persion中的name和student中的name同名的时候。
// 1. 如果本作用域能找到name，则对本作用于的name赋值。
// 2. 如果找不到，则查找继承的作用域，并进行赋值。
type persion struct {
	name string

}

type student struct {
	name string
	age int
	persion
}

func main() {
	var a student
	a.name = "yisan"
	a.age = 18
	// 如果要对person中的name进行定义，则需要显示的标记persion。
	a.persion.name = "ys"
	fmt.Println(a)
}
