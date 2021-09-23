package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {

}

type Cat struct {

}

func (dog *Dog) Speak() {
	fmt.Println("wang")
}

func (cat *Cat) Speak() {
	fmt.Println("miao")
}

func main() {
	dog := Dog{}
	cat := Cat{}

	animal := []Animal{&dog, &cat}

	for _, v := range animal {
		v.Speak()
	}
}