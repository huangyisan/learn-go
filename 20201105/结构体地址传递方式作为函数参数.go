package main

import "fmt"


type Student struct {
	name string
	id int
}

func changeName(student *Student, name string) {
	student.name = name
	fmt.Println("changeName: ", student)
}

func main() {
	// 该student为指针
	student := new(Student)
	student.name = "yisan"
	student.id = 32

	// 该student2为值
	student1 := Student{
		"huangys",
		32,
	}

	// 地址传递则改变了 main中的数据
	changeName(student,"huang")
	changeName(&student1,"huang")
	fmt.Println("main:" ,student)
	fmt.Println("main:" ,student1)


}