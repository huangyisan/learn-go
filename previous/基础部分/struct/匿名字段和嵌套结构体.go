package main

import "fmt"

// 结构体存在嵌套，且结构体定义了名称和类型就属于嵌套结构体。
// 结构体存在嵌套，且结构体只定义了类型，没有定义字段名，就属于匿名字段结构体。


type Human struct {
	name string
	age int
}

// Student属于嵌套结构体，因为human字段定义了名称以及类型，类型为Human
type Student struct {
	class string
	human Human
}

// Teacher属于匿名字段结构体，Human只有类型，没有定义字段名
type Teacher struct {
	lesson string
	Human
}

func main() {
	// 实例化student
	student := Student{
		"one",
		Human{
			"yisan",
			29,
		},
	}

	// 实例化teacher
	teacher := Teacher{
		"Chiness",
		Human{
			"yisan",
			29,
		},
	}
	fmt.Println(student)
	fmt.Println(teacher)

	//	修改信息
	student.human.name = "lisi"
	fmt.Println(student)
	fmt.Println(teacher.Human.name)

//	提升字段， teacher因为存在匿名字段，所以可以通过teacher直接访问匿名字段内的属性，就像是teacher直接继承了这些属性一样。
	fmt.Println(teacher.name)

}