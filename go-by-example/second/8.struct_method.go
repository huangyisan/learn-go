package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func (s *Student) printInfo () {
	fmt.Println(s.Name)
	fmt.Println(s.Age)
}

func main() {
	s := &Student{
		Name: "cc",
		Age: 29,
	}
	s.printInfo()
}