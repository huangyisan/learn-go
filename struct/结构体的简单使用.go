package main

import (
	"fmt"
)

//结构体类似，只包含init函数的python类

//声明方式
//type struct_variable_type struct {
//	member definition;
//	member definition;
//	...
//	member definition;
//}

// 定义名为books的结构体，包含books的一些属性

func main() {
type books struct {
	page int
	writer string
	color string
}

// 声明myBook为books类型
var myBook books

// 定义myBook属性
myBook.page=20
myBook.writer="yisan"
myBook.color="huang"

fmt.Println(myBook)

//声明中赋值
yourBook := books{
	20,
	"yisan",
	"huang",
}

fmt.Println(yourBook)
}