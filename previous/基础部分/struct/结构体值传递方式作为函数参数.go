package main

import "fmt"

type Student struct {
	name string
	id int
}

func changeName(student Student, name string) {
	student.name = name
	fmt.Println("changeName: ", student)
}

func main() {
	student := Student{
		"yisan",
		32,
	}

	// 因为是值传递方式,所以, 只在changeName函数中将name进行了改变,main中的student的name并未改变
	changeName(student,"huang")
	fmt.Println("main: ", student)


}